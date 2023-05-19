package server

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/environment/internal/store"
	"github.com/FoxFurry/scholarlabs/services/environment/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *ScholarLabsEnvironment) GetPublicPrototypes(ctx context.Context, req *emptypb.Empty) (*proto.GetPublicPrototypesResponse, error) {
	protos, err := p.service.GetPublicPrototypes(ctx)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to get public prototypes")
		return nil, err
	}

	return &proto.GetPublicPrototypesResponse{
		Prototypes: storeToProtoShortPrototypes(protos),
	}, nil
}

func storeToProtoShortPrototypes(storeProtos []store.PrototypeShort) []*proto.PrototypeShort {
	var protoProtos = make([]*proto.PrototypeShort, 0, len(storeProtos))

	for _, storeProto := range storeProtos {
		protoProtos = append(protoProtos, storeToProtoPrototypes(storeProto))
	}

	return protoProtos
}

func storeToProtoPrototypes(storeProto store.PrototypeShort) *proto.PrototypeShort {
	var protoProto = new(proto.PrototypeShort)

	protoProto.Name = storeProto.Name

	protoProto.UUID = storeProto.UUID
	protoProto.ShortDescription = storeProto.ShortDescription

	return protoProto
}
