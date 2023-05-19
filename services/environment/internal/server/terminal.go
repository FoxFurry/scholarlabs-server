package server

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/FoxFurry/scholarlabs/services/environment/proto"
	"github.com/FoxFurry/scholarlabs/virt"
)

const TERMINAL_INIT = "TERMINAL_INIT"

func (p *ScholarLabsEnvironment) BidirectionalTerminal(stream proto.Environment_BidirectionalTerminalServer) error {
	ctx := stream.Context()

	message, err := stream.Recv()
	if err != nil {
		if err == io.EOF {
			p.lg.WithField("msg", message).Info("client closed connection")
			return nil
		}

		p.lg.WithError(err).WithField("msg", message).Error("failed to receive stream")
		return err
	}

	if message.GetCommand() != TERMINAL_INIT {
		p.lg.WithField("msg", message).Info("init command not received")
		return fmt.Errorf("init command not received")
	}

	p.lg.WithField("msg", message).Info("init command received. starting terminal loop")

	env, err := p.service.GetEnvironmentByUUID(ctx, message.GetEnvironmentUUID())
	if err != nil {
		p.lg.WithError(err).WithField("msg", message).Error("failed to get environment")
		return fmt.Errorf("failed to get environment: %w", err)
	}

	prototype, err := p.service.GetPrototypeByUUID(ctx, env.PrototypeUUID)
	if err != nil {
		p.lg.WithError(err).WithField("msg", message).Error("failed to get prototype")
		return fmt.Errorf("failed to get prototype: %w", err)
	}

	terminalInstance, err := p.service.BidirectionalTerminal(ctx, prototype.Engine, env.MachineUUID)
	if err != nil {
		p.lg.WithError(err).WithField("msg", message).Error("failed to create terminal instance")
		return fmt.Errorf("failed to create terminal instance: %w", err)
	}

	p.lg.Infof("terminal: %+v", terminalInstance)
	p.lg.Infof("Conn: %+v", terminalInstance.GetConn())

	defer terminalInstance.Close()
	return p.terminalLoop(ctx, stream, terminalInstance)
}

func (p *ScholarLabsEnvironment) terminalLoop(ctx context.Context, stream proto.Environment_BidirectionalTerminalServer, terminal virt.Terminal) error {
	var (
		scanner = bufio.NewScanner(terminal.GetReader())
	)

	go func(sc *bufio.Scanner) {
		p.lg.Info("starting scan loop")

		for sc.Scan() {
			cmd := sc.Text()
			p.lg.WithField("req", cmd).Info("received response from terminal")
			err := stream.Send(&proto.BidirectionalTerminalResponse{Command: cmd})
			if err != nil {
				p.lg.WithError(err).WithField("req", cmd).Error("failed to receive response")

				return
			}
		}

		p.lg.Info("closing scan loop")

	}(scanner)

	for {
		select {
		case <-ctx.Done():
			p.lg.Info("context closed")

			return ctx.Err()
		default:
		}

		message, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				p.lg.WithField("req", message).Info("client closed connection")
				return nil
			}

			p.lg.WithError(err).WithField("req", message).Error("failed to receive stream")
			return err
		}

		p.lg.WithField("req", message).Info("received message")
		if _, err := terminal.GetConn().Write([]byte(message.GetCommand() + "\n")); err != nil {
			p.lg.WithError(err).WithField("req", message).Error("failed to send message to terminal")
			return err
		}
	}
}
