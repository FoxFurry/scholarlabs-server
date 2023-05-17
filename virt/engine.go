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

type Terminal interface {
	Close()
	GetConn() net.Conn
	GetReader() *bufio.Reader
}

type Engine interface {
	GetIdentifier(context.Context) string

	Spin(context.Context, string, string) (string, error)
	Destroy(context.Context, string) error

	GetDetails(context.Context, string) (*Details, error)
	StartTerminal(context.Context, string) (Terminal, error)
}
