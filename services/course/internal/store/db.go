package store

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/common/db"
	"github.com/FoxFurry/scholarlabs/services/course/internal/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DataStore interface {
	// User

	CreateUser(ctx context.Context, user User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUUID(ctx context.Context, userUUID string) (*User, error)

	// DB

	GetDB() *sqlx.DB
}

type store struct {
	sql *sqlx.DB
}

func (d *store) GetDB() *sqlx.DB { return d.sql }

func NewDataStore(cfg config.Config) (DataStore, error) {
	database, err := db.NewDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	if err != nil {
		return nil, err
	}

	return &store{
		sql: database,
	}, nil
}
