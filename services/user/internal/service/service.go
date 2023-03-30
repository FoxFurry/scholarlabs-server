package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/FoxFurry/scholarlabs/services/user/internal/httperr"
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
			return httperr.New(fmt.Sprintf("%s: entry already exists", msg), http.StatusBadRequest)
		case 1741:
			return httperr.New(fmt.Sprintf("%s: key not found", msg), http.StatusNotFound)
		}
	}
	// TODO: Change in live environment
	return httperr.New(fmt.Sprintf("%s: unknown internal error: %v", msg, err), http.StatusInternalServerError)
}
