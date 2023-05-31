package virt

import (
	"bufio"
	"context"
	"net"
)

type Details struct {
	IsRunning bool
	IP        string
}

type PrototypeData struct {
	EngineRef string
	Env       []string
	Cmd       []string
}

type Terminal interface {
	Close()
	GetConn() net.Conn
	GetReader() *bufio.Reader
}

type Engine interface {
	GetIdentifier(context.Context) string

	Spin(context.Context, PrototypeData) (string, error)
	Destroy(context.Context, string) error

	GetDetails(context.Context, string) (*Details, error)
	StartTerminal(context.Context, string, []string) (Terminal, error)
}
