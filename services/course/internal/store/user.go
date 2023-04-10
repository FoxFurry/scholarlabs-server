package store

import (
	"context"
	"time"
)

type User struct {
	ID       uint64
	UUID     string
	Email    string
	Username string
	Password string

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (d *store) CreateUser(ctx context.Context, user User) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO users (email, username, password, uuid) VALUES (?, ?, ?, ?)`, user.Email, user.Username, user.Password, user.UUID)
	return err
}

func (d *store) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var u User

	if err := d.sql.GetContext(ctx, &u, `SELECT * FROM users WHERE email=?`, email); err != nil {
		return nil, err
	}

	return &u, nil
}

func (d *store) GetUserByUUID(ctx context.Context, UUID string) (*User, error) {
	var u User

	if err := d.sql.GetContext(ctx, &u, `SELECT * FROM users WHERE uuid=?`, UUID); err != nil {
		return nil, err
	}

	return &u, nil
}
