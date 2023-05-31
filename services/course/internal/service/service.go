package service

import (
	"context"
	"fmt"

	"github.com/FoxFurry/scholarlabs/services/course/internal/store"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-sql-driver/mysql"
)

var bucketName = "scholarlabs"

type Service interface {
	// Course
	Enroll(ctx context.Context, userUUID, courseUUID string) error
	Unenroll(ctx context.Context, userUUID, courseUUID string) error

	GetEnrolledCoursesForUser(ctx context.Context, userUUID string) ([]store.Course, error)

	CreateCourse(ctx context.Context, c store.Course, thumbnail, background []byte) (string, error)

	GetCourseInfo(ctx context.Context, courseUUID string) (*store.Course, error)
	//GetCourseDashBoard(ctx context.Context, courseUUID string) (*store.CourseDashboard, error)
	GetAllPublicCourses(ctx context.Context) ([]store.Course, error)

	// Page
	CreatePage(ctx context.Context, p store.Page) error
	DeletePage(ctx context.Context, courseUUID, pageUUID string) error
	GetPageByID(ctx context.Context, courseUUID, pageUUID string) (*store.Page, error)
	GetCourseToc(ctx context.Context, courseUUID string) ([]store.PageIdentifier, error)
}

type service struct {
	db       store.DataStore
	s3Client *s3.S3
}

func New(datastore store.DataStore, bucket *s3.S3) Service {
	//if err := checkOrCreateBucket(bucket, bucketName); err != nil {
	//	panic(err)
	//}

	return &service{
		db:       datastore,
		s3Client: bucket,
	}
}

func checkOrCreateBucket(bucket *s3.S3, bucketName string) error {
	_, err := bucket.HeadBucket(&s3.HeadBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		_, err = bucket.CreateBucket(&s3.CreateBucketInput{
			Bucket: &bucketName,
		})
		if err != nil {
			return fmt.Errorf("could not create bucket: %v", err)
		}
	}
	return nil
}

func handleDBError(err error, msg string) error {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case 1062:
			return fmt.Errorf("%s: entry already exists", msg)
		case 1741:
			return fmt.Errorf("%s: key not found", msg)
		}
	}
	// TODO: Change in live environment
	return fmt.Errorf("%s: unknown internal error: %v", msg, err)
}
