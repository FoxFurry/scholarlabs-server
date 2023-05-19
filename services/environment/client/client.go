package client

import (
	"github.com/FoxFurry/scholarlabs/services/environment/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewEnvironmentClient(address string) (proto.EnvironmentClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return proto.NewEnvironmentClient(conn), nil
}
