package store

import (
	"context"
	"time"
)

type Course struct {
	ID               uint64
	UUID             string
	AuthorUUID       string `db:"author_uuid"`
	Title            string
	ShortDescription string `db:"short_description"`
	Description      string
	Thumbnail        string
	Background       string

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CourseDashboard struct {
	Course

	AssignmentsProgress []struct {
		AssignmentUUID string `db:"assignment_uuid"`
	}
}

func (d *store) CreateCourse(ctx context.Context, course Course) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO courses (uuid, author_uuid, title, short_description, description, thumbnail, background) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		course.UUID,
		course.AuthorUUID,
		course.Title,
		course.ShortDescription,
		course.Description,
		course.Thumbnail,
		course.Background,
	)
	return err
}

func (d *store) GetCourseByUUID(ctx context.Context, uuid string) (*Course, error) {
	var u Course

	if err := d.sql.GetContext(ctx, &u, `SELECT * FROM courses WHERE uuid=?`, uuid); err != nil {
		return nil, err
	}

	return &u, nil
}

func (d *store) GetPublicCourses(ctx context.Context) ([]Course, error) {
	rows, err := d.sql.QueryContext(ctx, `SELECT * from courses`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		courses []Course
		buffer  Course
	)

	for rows.Next() {
		if err := rows.Scan(
			&buffer.ID,
			&buffer.UUID,
			&buffer.AuthorUUID,
			&buffer.Title,
			&buffer.ShortDescription,
			&buffer.Description,
			&buffer.Thumbnail,
			&buffer.Background,
			&buffer.CreatedAt,
			&buffer.UpdatedAt,
		); err != nil {
			d.lg.WithError(err).Error("failed to read course from db")
			return nil, err
		}

		courses = append(courses, buffer)
	}

	return courses, nil
}

func (d *store) GetEnrolledCoursesForUser(ctx context.Context, userUUID string) ([]Course, error) {
	rows, err := d.sql.QueryContext(ctx, `SELECT uuid, author_uuid, title, short_description, thumbnail FROM courses INNER JOIN enrolls ON enrolls.course_uuid = uuid WHERE enrolls.user_uuid = ?`,
		userUUID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		courses []Course
		buffer  Course
	)

	for rows.Next() {
		if err := rows.Scan(
			&buffer.UUID,
			&buffer.AuthorUUID,
			&buffer.Title,
			&buffer.ShortDescription,
			&buffer.Thumbnail,
		); err != nil {
			d.lg.WithError(err).Error("failed to read course from db")
		}

		courses = append(courses, buffer)
	}

	return courses, nil
}
