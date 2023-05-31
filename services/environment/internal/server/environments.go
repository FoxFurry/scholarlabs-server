package server

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/environment/internal/store"
	"github.com/FoxFurry/scholarlabs/services/environment/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *ScholarLabsEnvironment) CreateEnvironment(ctx context.Context, req *proto.CreateEnvironmentRequest) (*emptypb.Empty, error) {
	if _, err := p.service.CreateEnvironment(ctx, store.Environment{
		Name:          req.GetName(),
		OwnerUUID:     req.GetOwnerUUID(),
		PrototypeUUID: req.GetPrototypeUUID(),
	}); err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to create an environment")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (p *ScholarLabsEnvironment) GetEnvironmentsForUser(ctx context.Context, req *proto.GetEnvironmentsForUserRequest) (*proto.GetEnvironmentsForUserResponse, error) {
	courses, err := p.service.GetEnvironmentsForUser(ctx, req.OwnerUUID)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to get environments for user")
		return nil, err
	}

	return &proto.GetEnvironmentsForUserResponse{
		Environments: storeToProtoShortEnvs(courses),
	}, nil
}

func (p *ScholarLabsEnvironment) GetEnvironmentDetails(ctx context.Context, req *proto.GetEnvironmentDetailsRequest) (*proto.GetEnvironmentDetailsResponse, error) {
	env, err := p.service.GetEnvironmentByUUID(ctx, req.EnvironmentUUID)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to get environment details")
		return nil, err
	}

	return &proto.GetEnvironmentDetailsResponse{
		Environment: storeToProtoFullEnv(env),
	}, nil
}

func storeToProtoShortEnvs(storeEnvs []store.Environment) []*proto.EnvironmentShort {
	var protoEnvs = make([]*proto.EnvironmentShort, 0, len(storeEnvs))

	for _, storeEnv := range storeEnvs {
		protoEnvs = append(protoEnvs, storeToProtoShortEnv(storeEnv))
	}

	return protoEnvs
}

func storeToProtoShortEnv(storeEnv store.Environment) *proto.EnvironmentShort {
	var protoEnv = new(proto.EnvironmentShort)

	protoEnv.Name = storeEnv.Name

	protoEnv.UUID = storeEnv.UUID
	protoEnv.OwnerUUID = storeEnv.OwnerUUID

	return protoEnv
}

func storeToProtoFullEnv(storeEnv *store.Environment) *proto.EnvironmentFull {
	var protoEnv = new(proto.EnvironmentFull)
	protoEnv.Short = new(proto.EnvironmentShort)

	protoEnv.Short.Name = storeEnv.Name
	protoEnv.Short.UUID = storeEnv.UUID
	protoEnv.Short.OwnerUUID = storeEnv.OwnerUUID
	protoEnv.Short.PrototypeUUID = storeEnv.PrototypeUUID

	protoEnv.MachineUUID = storeEnv.MachineUUID

	return protoEnv
}
