package service

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/course/internal/store"
	"github.com/google/uuid"
)

func (p *service) CreateCourse(ctx context.Context, c store.Course) (*store.Course, error) {
	c.UUID = uuid.New().String()

	if err := p.db.CreateCourse(ctx, c); err != nil {
		return nil, handleDBError(err, "could not create course")
	}

	return &c, nil
}

func (p *service) GetAllPublicCourses(ctx context.Context) ([]store.Course, error) {
	courses, err := p.db.GetPublicCourses(ctx)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (p *service) GetEnrolledCoursesForUser(ctx context.Context, userUUID string) ([]store.Course, error) {
	courses, err := p.db.GetEnrolledCoursesForUser(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (p *service) GetCourseInfo(ctx context.Context, courseUUID string) (*store.Course, error) {
	course, err := p.db.GetCourseByUUID(ctx, courseUUID)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (p *service) Enroll(ctx context.Context, userUUID, courseUUID string) error {
	return p.db.Enroll(ctx, userUUID, courseUUID)
}

func (p *service) Unenroll(ctx context.Context, userUUID, courseUUID string) error {
	return p.db.Unenroll(ctx, userUUID, courseUUID)
}
