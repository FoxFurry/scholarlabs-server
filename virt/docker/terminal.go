package docker

import (
	"bufio"
	"net"

	"github.com/docker/docker/api/types"
)

type dockerTerminalAdapter struct {
	term *types.HijackedResponse
}

func (d *dockerTerminalAdapter) Close() {
	d.Close()
}

func (d *dockerTerminalAdapter) GetConn() net.Conn {
	return d.term.Conn
}

func (d *dockerTerminalAdapter) GetReader() *bufio.Reader {
	return d.term.Reader
}
