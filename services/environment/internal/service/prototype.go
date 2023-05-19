package service

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/environment/internal/store"
)

func (s *service) GetPublicPrototypes(ctx context.Context) ([]store.PrototypeShort, error) {
	protos, err := s.db.GetPublicPrototypes(ctx)
	if err != nil {
		return nil, err
	}

	return protos, nil
}

func (s *service) GetPrototypeByUUID(ctx context.Context, protoUUID string) (*store.PrototypeFull, error) {
	proto, err := s.db.GetPrototypeByUUID(ctx, protoUUID)
	if err != nil {
		return nil, err
	}

	return proto, nil
}
