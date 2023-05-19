package store

import (
	"context"
	"time"
)

type Environment struct {
	UUID          string
	Name          string
	OwnerUUID     string
	PrototypeUUID string
	MachineUUID   string

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (d *store) CreateEnvironment(ctx context.Context, env Environment) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO environments (uuid, name, owner_uuid, prototype_uuid, machine_uuid) VALUES (?, ?, ?, ?, ?)`,
		env.UUID,
		env.Name,
		env.OwnerUUID,
		env.PrototypeUUID,
		env.MachineUUID,
	)
	return err
}

func (d *store) GetEnvironmentsForUser(ctx context.Context, ownerUUID string) ([]Environment, error) {
	rows, err := d.sql.QueryContext(ctx, `SELECT uuid, name, owner_uuid, prototype_uuid, machine_uuid FROM environments WHERE owner_uuid = ?`,
		ownerUUID,
	)
	if err != nil {
		return nil, err
	}

	var (
		envs   []Environment
		buffer Environment
	)

	for rows.Next() {
		if err := rows.Scan(
			&buffer.UUID,
			&buffer.Name,
			&buffer.OwnerUUID,
			&buffer.PrototypeUUID,
			&buffer.MachineUUID,
		); err != nil {
			d.lg.WithError(err).Error("failed to read environments from db")
		}

		envs = append(envs, buffer)
	}

	return envs, nil
}

func (d *store) GetEnvironmentDetails(ctx context.Context, envUUID string) (*Environment, error) {
	row := d.sql.QueryRowContext(ctx, `SELECT uuid, name, owner_uuid, prototype_uuid, machine_uuid FROM environments WHERE uuid = ?`,
		envUUID,
	)

	var env Environment
	if err := row.Scan(
		&env.UUID,
		&env.Name,
		&env.OwnerUUID,
		&env.PrototypeUUID,
		&env.MachineUUID,
	); err != nil {
		return nil, err
	}

	return &env, nil
}
