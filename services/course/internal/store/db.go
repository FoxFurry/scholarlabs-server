package store

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/common/db"
	"github.com/FoxFurry/scholarlabs/services/course/internal/config"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DataStore interface {

	// Course
	CreateCourse(context.Context, Course) error

	GetCourseByUUID(context.Context, string) (*Course, error)
	GetPublicCourses(context.Context) ([]Course, error)
	GetEnrolledCoursesForUser(context.Context, string) ([]Course, error)

	// Enrolls

	Enroll(context.Context, string, string) error
	Unenroll(context.Context, string, string) error

	// Page
	CreatePage(context.Context, Page) error
	DeletePage(context.Context, string, string) error
	GetPageByID(context.Context, string, string) (*Page, error)
	GetCourseToC(context.Context, string) ([]PageIdentifier, error)
	GetNumberOfPagesForCourse(context.Context, string) (int, error)
	//
	GetAssignmentsProgressForUser(context.Context, string, string) ([]AssignmentProgress, error)
	StartAssignment(context.Context, string, string, string) error
	IsAssignmentStarted(context.Context, string, string, string) (bool, error)

	// DB
	GetDB() *sqlx.DB
}

type store struct {
	sql *sqlx.DB
	lg  *logrus.Logger
}

func (d *store) GetDB() *sqlx.DB { return d.sql }

func NewDataStore(cfg config.Config, logger *logrus.Logger) (DataStore, error) {
	database, err := db.NewDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	if err != nil {
		return nil, err
	}

	return &store{
		sql: database,
		lg:  logger,
	}, nil
}
