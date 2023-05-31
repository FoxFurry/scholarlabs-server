package docker

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/FoxFurry/scholarlabs/virt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const identity = "docker"

type docker struct {
	cli *client.Client
}

func New(ctx context.Context) (virt.Engine, error) {
	cli, err := client.NewClientWithOpts(client.WithHost("tcp://165.22.23.225:2375")) // client.WithHost("tcp://165.22.67.123:2375")
	if err != nil {
		return nil, err
	}

	cli.NegotiateAPIVersion(ctx)

	go func(cli *client.Client) {
		<-ctx.Done()

		if err := cli.Close(); err != nil {
			panic(err)
		}
	}(cli)

	return &docker{
		cli: cli,
	}, nil
}

func (d *docker) GetIdentifier(ctx context.Context) string {
	return identity
}

func (d *docker) Spin(ctx context.Context, data virt.PrototypeData) (string, error) {
	pull, err := d.cli.ImagePull(ctx, data.EngineRef, types.ImagePullOptions{})
	if err != nil {
		return "", fmt.Errorf("could not pull image: %w", err)
	}

	defer func(pull io.ReadCloser) {
		if err := pull.Close(); err != nil {
			panic(err)
		}
	}(pull)

	resp, err := d.cli.ContainerCreate(ctx, &container.Config{
		Image: data.EngineRef,
		Cmd:   data.Cmd,
		Env:   data.Env,
	}, nil, nil, nil, "")
	if err != nil {
		return "", fmt.Errorf("could not create container: %w", err)
	}

	time.Sleep(5 * time.Second)

	if err := d.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}

	return resp.ID, nil
}

func (d *docker) Destroy(ctx context.Context, machineIdentifier string) error {
	if err := d.cli.ContainerStop(ctx, machineIdentifier, container.StopOptions{}); err != nil {
		return fmt.Errorf("failed to stop container: %w", err)
	}

	return d.cli.ContainerRemove(ctx, machineIdentifier, types.ContainerRemoveOptions{
		RemoveLinks:   true,
		RemoveVolumes: true,
		Force:         true,
	})
}

func (d *docker) GetDetails(ctx context.Context, machineIdentifier string) (*virt.Details, error) {
	resp, err := d.cli.ContainerInspect(ctx, machineIdentifier)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect container: %w", err)
	}

	fmt.Println(resp.NetworkSettings)

	return &virt.Details{}, nil
}

func (d *docker) StartTerminal(ctx context.Context, machineIdentifier string, execCmd []string) (virt.Terminal, error) {
	execPoint, err := d.cli.ContainerExecCreate(ctx, machineIdentifier, types.ExecConfig{
		AttachStderr: true,
		AttachStdout: true,
		AttachStdin:  true,
		Tty:          false,
		Cmd:          execCmd,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create exec point: %w", err)
	}

	terminalHijack, err := d.cli.ContainerExecAttach(ctx, execPoint.ID, types.ExecStartCheck{})
	if err != nil {
		return nil, fmt.Errorf("failed to attach to the exec point: %w", err)
	}

	return &dockerTerminalAdapter{
		term: &terminalHijack,
	}, nil
}
