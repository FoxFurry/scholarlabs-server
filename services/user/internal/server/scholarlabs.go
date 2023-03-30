package server

import (
	"github.com/FoxFurry/scholarlabs/services/user/internal/config"
	"github.com/FoxFurry/scholarlabs/services/user/internal/service"
	"github.com/FoxFurry/scholarlabs/services/user/internal/store"
	"github.com/FoxFurry/scholarlabs/services/user/internal/util"
	"github.com/FoxFurry/scholarlabs/services/user/proto"
	"github.com/sirupsen/logrus"
)

const (
	issuer      = "scholarlabs"
	tokenSecret = "5dd0bf305c1eb5b832dbc4169c84ba0aa51704da74b8d2e953dca7b276ee8b0c821e8a764f16fd183c50ca9d9b655cf6159564a1554da81ee16fe01866a462e225ad779b472a62b15d2861c54579875709da3e025e916ab3ac89b165359d0ac529e3739a513eb0de1a2350ab9f741"
)

type ScholarLabs struct {
	service service.Service
	jwt     util.JWTProvider
	cfg     config.Config
	lg      *logrus.Logger
	proto.UnimplementedUserServer
}

func New(cfg config.Config, ds store.DataStore, logger *logrus.Logger) (*ScholarLabs, error) {
	return &ScholarLabs{
		service: service.New(ds),
		jwt:     util.NewJWT(),
		cfg:     cfg,
		lg:      logger,
	}, nil
}
