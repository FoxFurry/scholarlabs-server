package server

import (
	"github.com/FoxFurry/scholarlabs/services/course/internal/config"
	"github.com/FoxFurry/scholarlabs/services/course/internal/service"
	"github.com/FoxFurry/scholarlabs/services/course/internal/store"
	"github.com/FoxFurry/scholarlabs/services/course/proto"
	"github.com/sirupsen/logrus"
)

type ScholarLabs struct {
	service service.Service
	cfg     config.Config
	lg      *logrus.Logger
	proto.UnimplementedCoursesServer
}

func New(cfg config.Config, ds store.DataStore, logger *logrus.Logger) (*ScholarLabs, error) {
	return &ScholarLabs{
		service: service.New(ds),
		cfg:     cfg,
		lg:      logger,
	}, nil
}
