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
	CreateCourse(ctx context.Context, course Course) error

	GetCourseByUUID(ctx context.Context, uuid string) (*Course, error)
	GetPublicCourses(ctx context.Context) ([]Course, error)
	GetEnrolledCoursesForUser(ctx context.Context, userUUID string) ([]Course, error)

	// Enrolls

	Enroll(ctx context.Context, userUUID, courseUUID string) error
	Unenroll(ctx context.Context, userUUID, courseUUID string) error

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
