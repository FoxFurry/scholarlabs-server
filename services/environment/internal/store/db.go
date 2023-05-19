package store

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/common/db"
	"github.com/FoxFurry/scholarlabs/services/environment/internal/config"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DataStore interface {

	// Env
	CreateEnvironment(context.Context, Environment) error
	GetEnvironmentsForUser(context.Context, string) ([]Environment, error)
	GetEnvironmentDetails(context.Context, string) (*Environment, error)

	// Prototype
	GetPublicPrototypes(ctx context.Context) ([]PrototypeShort, error)
	GetPrototypeByUUID(context.Context, string) (*PrototypeFull, error)
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
