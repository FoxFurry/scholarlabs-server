package server

import (
	"github.com/FoxFurry/scholarlabs/services/course/internal/config"
	"github.com/FoxFurry/scholarlabs/services/course/internal/service"
	"github.com/FoxFurry/scholarlabs/services/course/internal/store"
	"github.com/FoxFurry/scholarlabs/services/course/proto"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
)

type ScholarLabs struct {
	service service.Service
	cfg     config.Config
	lg      *logrus.Logger
	proto.UnimplementedCoursesServer
}

func New(cfg config.Config, ds store.DataStore, bucket *s3.S3, logger *logrus.Logger) (*ScholarLabs, error) {
	return &ScholarLabs{
		service: service.New(ds, bucket),
		cfg:     cfg,
		lg:      logger,
	}, nil
}
