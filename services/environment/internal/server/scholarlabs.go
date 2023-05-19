package server

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/environment/internal/config"
	"github.com/FoxFurry/scholarlabs/services/environment/internal/service"
	"github.com/FoxFurry/scholarlabs/services/environment/internal/store"
	"github.com/FoxFurry/scholarlabs/services/environment/proto"
	"github.com/sirupsen/logrus"
)

type ScholarLabsEnvironment struct {
	service service.Service
	cfg     config.Config
	lg      *logrus.Logger
	proto.UnimplementedEnvironmentServer
}

func New(ctx context.Context, cfg config.Config, ds store.DataStore, logger *logrus.Logger) (*ScholarLabsEnvironment, error) {
	return &ScholarLabsEnvironment{
		service: service.New(ctx, ds),
		cfg:     cfg,
		lg:      logger,
	}, nil
}
