package service

import (
	"context"
	"fmt"
	"log"

	"github.com/FoxFurry/scholarlabs/services/environment/internal/store"
	"github.com/FoxFurry/scholarlabs/virt"
	"github.com/FoxFurry/scholarlabs/virt/docker"
	"github.com/go-sql-driver/mysql"
)

type Service interface {
	// Course

	CreateEnvironment(context.Context, store.Environment) (*store.Environment, error)
	GetEnvironmentsForUser(context.Context, string) ([]store.Environment, error)
	GetEnvironmentByUUID(context.Context, string) (*store.Environment, error)

	GetPublicPrototypes(context.Context) ([]store.PrototypeShort, error)
	GetPrototypeByUUID(context.Context, string) (*store.PrototypeFull, error)

	BidirectionalTerminal(ctx context.Context, engine, termRef string) (virt.Terminal, error)

	CreateRoom(ctx context.Context, environmentUUID string) (string, error)
	ResolveRoom(ctx context.Context, room string) (string, error)
}

type service struct {
	db           store.DataStore
	dockerEngine virt.Engine
}

func New(ctx context.Context, datastore store.DataStore) Service {
	dkEng, err := docker.New(ctx)
	if err != nil {
		log.Fatalf("could not create docker engine: %v", err)
	}

	return &service{
		db:           datastore,
		dockerEngine: dkEng,
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
