package store

import (
	"context"
)

type AssignmentProgress struct {
	AssignmentUUID  string `db:"assignment_uuid"`
	UserUUID        string `db:"user_uuid"`
	EnvironmentUUID string `db:"environment_uuid"`
	Grade           uint64
}

func (d *store) GetAssignmentsProgressForUser(ctx context.Context, userUUID string, courseUUID string) ([]AssignmentProgress, error) {
	rows, err := d.sql.QueryContext(ctx, `SELECT * FROM assignments_progress WHERE user_uuid=? and course_uuid=?`, userUUID, courseUUID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		assignmentsProgress []AssignmentProgress
		buffer              AssignmentProgress
	)

	for rows.Next() {
		if err := rows.Scan(
			&buffer.AssignmentUUID,
			&buffer.UserUUID,
			&buffer.EnvironmentUUID,
			&buffer.Grade,
		); err != nil {
			return nil, err
		}

		assignmentsProgress = append(assignmentsProgress, buffer)
	}

	return assignmentsProgress, nil
}

func (d *store) StartAssignment(ctx context.Context, userUUID string, courseUUID string, assignmentUUID string) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO assignments_progress (assignment_uuid, user_uuid, environment_uuid) VALUES (?, ?, ?)`,
		assignmentUUID,
		userUUID,
		courseUUID,
	)
	return err
}

func (d *store) IsAssignmentStarted(context.Context, string, string, string) (bool, error) {
	return false, nil
}
