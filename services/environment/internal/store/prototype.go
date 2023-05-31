package store

import (
	"context"
	"time"
)

type PrototypeShort struct {
	UUID             string
	Name             string
	ShortDescription string `db:"short_description"`
	Engine           string
}

type PrototypeFull struct {
	PrototypeShort

	FullDescription string `db:"full_description"`
	Engine          string
	EngineRef       string `db:"engine_ref"`

	Env       string
	CmdTailer string    `db:"cmd_tailer"`
	Cmd       string    `db:"cmd"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (d *store) GetPublicPrototypes(ctx context.Context) ([]PrototypeShort, error) {
	rows, err := d.sql.QueryContext(ctx, `SELECT uuid, name, short_description, engine FROM prototypes`)
	if err != nil {
		return nil, err
	}

	var (
		protos []PrototypeShort
		buffer PrototypeShort
	)

	for rows.Next() {
		if err := rows.Scan(
			&buffer.UUID,
			&buffer.Name,
			&buffer.ShortDescription,
			&buffer.Engine,
		); err != nil {
			d.lg.WithError(err).Error("failed to read environments from db")
		}

		protos = append(protos, buffer)
	}

	return protos, nil
}

func (d *store) GetPrototypeByUUID(ctx context.Context, uuid string) (*PrototypeFull, error) {
	var proto PrototypeFull

	if err := d.sql.QueryRowContext(ctx, `SELECT uuid, name, short_description, full_description, engine, engine_ref, env, cmd_tailer, cmd, created_at, updated_at FROM prototypes WHERE uuid = ?`, uuid).Scan(
		&proto.UUID,
		&proto.Name,
		&proto.ShortDescription,
		&proto.FullDescription,
		&proto.Engine,
		&proto.EngineRef,
		&proto.Env,
		&proto.CmdTailer,
		&proto.Cmd,
		&proto.CreatedAt,
		&proto.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &proto, nil
}
