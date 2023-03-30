package service

import (
	"context"
	"fmt"

	"github.com/FoxFurry/scholarlabs/services/common/hash"
	"github.com/FoxFurry/scholarlabs/services/user/internal/store"
	"github.com/google/uuid"
)

func (p *service) CreateNewUser(ctx context.Context, u store.User) (*store.User, error) {
	u.UUID = uuid.New().String()
	hashedPassword, err := hash.NewPassword(u.Password)
	if err != nil {
		return nil, fmt.Errorf("could not hash password: %w", err)
	}

	u.Password = hashedPassword

	if err = p.db.CreateUser(ctx, u); err != nil {
		return nil, handleDBError(err, "could not create user")
	}

	return &u, nil
}

func (p *service) GetUserByEmail(ctx context.Context, email string) (*store.User, error) {
	user, err := p.db.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, handleDBError(err, "could not get user by email")
	}

	return user, nil
}

func (p *service) GetUserByUUID(ctx context.Context, userUUID string) (*store.User, error) {
	user, err := p.db.GetUserByUUID(ctx, userUUID)
	if err != nil {
		return nil, handleDBError(err, "could not get user by uuid")
	}

	return user, nil
}
