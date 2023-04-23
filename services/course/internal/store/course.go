package store

import (
	"context"
	"time"
)

type Course struct {
	ID          uint64
	UUID        string
	AuthorUUID  string
	Title       string
	Description string
	Thumbnail   string

	Text string

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (d *store) CreateCourse(ctx context.Context, course Course) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO courses (uuid, author_uuid, title, description, thumbnail, text) VALUES (?, ?, ?, ?, ?, ?)`,
		course.UUID,
		course.AuthorUUID,
		course.Title,
		course.Description,
		course.Thumbnail,
		course.Text,
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
			&buffer.Description,
			&buffer.Thumbnail,
			&buffer.Text,
			&buffer.CreatedAt,
			&buffer.UpdatedAt,
		); err != nil {
			d.lg.WithError(err).Error("failed to read course from db")
		}

		courses = append(courses, buffer)
	}

	return courses, nil
}

func (d *store) GetEnrolledCoursesForUser(ctx context.Context, userUUID string) ([]Course, error) {
	rows, err := d.sql.QueryContext(ctx, `SELECT uuid, author_uuid, title, description, thumbnail FROM courses INNER JOIN enrolls ON enrolls.course_uuid = uuid WHERE enrolls.user_uuid = ?`,
		userUUID,
	)
	if err != nil {
		return nil, err
	}

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
			&buffer.Description,
			&buffer.Thumbnail,
			&buffer.CreatedAt,
			&buffer.UpdatedAt,
		); err != nil {
			d.lg.WithError(err).Error("failed to read course from db")
		}

		courses = append(courses, buffer)
	}

	return courses, nil
}

func (d *store) Enroll(ctx context.Context, userUUID, courseUUID string) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO enrolls (user_uuid, course_uuid) VALUES (?, ?)`,
		userUUID,
		courseUUID,
	)
	return err
}

func (d *store) Unenroll(ctx context.Context, userUUID, courseUUID string) error {
	_, err := d.sql.ExecContext(ctx, `DELETE FROM enrolls WHERE user_uuid=? and course_uuid=?`,
		userUUID,
		courseUUID,
	)
	return err
}
