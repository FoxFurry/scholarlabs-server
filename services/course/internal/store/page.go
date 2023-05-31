package store

import (
	"context"
	"time"
)

type PageType string

var (
	PageTypeLesson     PageType = "LESSON"
	PageTypeAssignment PageType = "ASSIGNMENT"
)

type PageIdentifier struct {
	PageID uint64 `db:"page_id"`
	Title  string
	Type   PageType
}

type Page struct {
	ID         uint64
	CourseUUID string `db:"course_uuid"`
	Data       []byte

	PageIdentifier

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (d *store) CreatePage(ctx context.Context, page Page) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO pages (page_id, title, course_uuid, type, data) VALUES (?, ?, ?, ?, ?)`,
		page.PageID,
		page.Title,
		page.CourseUUID,
		page.Type,
		page.Data,
	)
	return err
}

func (d *store) DeletePage(ctx context.Context, courseUUID, pageID string) error {
	_, err := d.sql.ExecContext(ctx, `DELETE FROM pages WHERE course_uuid=? AND page_id=?`, courseUUID, pageID)
	return err
}

func (d *store) GetPageByID(ctx context.Context, courseUUID, pageID string) (*Page, error) {
	var u Page

	if err := d.sql.GetContext(ctx, &u, `SELECT * FROM pages WHERE course_uuid=? AND page_id=?`, courseUUID, pageID); err != nil {
		return nil, err
	}

	return &u, nil
}

func (d *store) GetCourseToC(ctx context.Context, courseUUID string) ([]PageIdentifier, error) {
	rows, err := d.sql.QueryContext(ctx, `SELECT page_id, title, type FROM pages WHERE course_uuid=?`, courseUUID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		pages  []PageIdentifier
		buffer PageIdentifier
	)

	for rows.Next() {
		if err := rows.Scan(
			&buffer.PageID,
			&buffer.Title,
			&buffer.Type,
		); err != nil {
			return nil, err
		}

		pages = append(pages, buffer)
	}

	return pages, nil
}

func (d *store) GetNumberOfPagesForCourse(ctx context.Context, courseUUID string) (int, error) {
	var count int
	err := d.sql.GetContext(ctx, &count, `SELECT COUNT(*) FROM pages WHERE course_uuid=?`, courseUUID)
	if err != nil {
		return 0, err
	}

	return count, nil
}
