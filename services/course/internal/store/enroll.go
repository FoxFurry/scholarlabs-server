package store

import (
	"context"
)

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
