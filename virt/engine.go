package virt

import (
	"context"

	"github.com/docker/docker/api/types"
)

type Details struct {
	IsRunning bool
	IP        string
}

type Engine interface {
	GetIdentifier(context.Context) string

	Spin(context.Context, string, string) (string, error)
	Destroy(context.Context, string) error

	GetDetails(context.Context, string) (*Details, error)
	StartTerminal(context.Context, string) (*types.HijackedResponse, error)
}
