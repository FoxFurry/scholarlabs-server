package service

import (
	"context"
	"fmt"

	"github.com/FoxFurry/scholarlabs/services/user/internal/store"
	"github.com/go-sql-driver/mysql"
)

type Service interface {
	// User

	CreateNewUser(ctx context.Context, u store.User) (*store.User, error)
	GetUserByEmail(ctx context.Context, mail string) (*store.User, error)
	GetUserByUUID(ctx context.Context, userUUID string) (*store.User, error)
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
