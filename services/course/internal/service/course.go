package service

import (
	"bytes"
	"context"

	"github.com/FoxFurry/scholarlabs/services/course/internal/store"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func (p *service) CreateCourse(ctx context.Context, c store.Course, thumbnail, background []byte) (string, error) {
	c.UUID = uuid.New().String()

	//if _, err := p.s3Client.PutObject(createObject(fmt.Sprintf("course/%s/thumbnail.png", c.UUID), thumbnail)); err != nil {
	//	return nil, err
	//}
	//
	//if _, err := p.s3Client.PutObject(createObject(fmt.Sprintf("course/%s/background.png", c.UUID), background)); err != nil {
	//	return nil, err
	//}

	if err := p.db.CreateCourse(ctx, c); err != nil {
		return "", handleDBError(err, "could not create course")
	}

	return c.UUID, nil
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

func createObject(filename string, content []byte) *s3.PutObjectInput {
	return &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &filename,
		Body:   bytes.NewReader(content),
		ACL:    aws.String("public-read"),
	}
}
