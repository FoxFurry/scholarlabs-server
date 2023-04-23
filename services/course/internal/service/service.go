package service

import (
	"context"
	"fmt"

	"github.com/FoxFurry/scholarlabs/services/course/internal/store"
	"github.com/go-sql-driver/mysql"
)

type Service interface {
	// Course
	Enroll(ctx context.Context, userUUID, courseUUID string) error
	Unenroll(ctx context.Context, userUUID, courseUUID string) error

	GetEnrolledCoursesForUser(ctx context.Context, userUUID string) ([]store.Course, error)

	CreateCourse(ctx context.Context, c store.Course) (*store.Course, error)

	GetCourseInfo(ctx context.Context, courseUUID string) (*store.Course, error)
	GetAllPublicCourses(ctx context.Context) ([]store.Course, error)
}

type service struct {
	db store.DataStore
}

func New(datastore store.DataStore) Service {
	return &service{
		db: datastore,
	}
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
