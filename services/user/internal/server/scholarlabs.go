package server

import (
	"context"
	"fmt"
	"time"

	"github.com/FoxFurry/scholarlabs/services/common/hash"
	"github.com/FoxFurry/scholarlabs/services/user/internal/config"
	"github.com/FoxFurry/scholarlabs/services/user/internal/service"
	"github.com/FoxFurry/scholarlabs/services/user/internal/store"
	"github.com/FoxFurry/scholarlabs/services/user/internal/util"
	"github.com/FoxFurry/scholarlabs/services/user/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	issuer      = "scholarlabs"
	tokenSecret = "5dd0bf305c1eb5b832dbc4169c84ba0aa51704da74b8d2e953dca7b276ee8b0c821e8a764f16fd183c50ca9d9b655cf6159564a1554da81ee16fe01866a462e225ad779b472a62b15d2861c54579875709da3e025e916ab3ac89b165359d0ac529e3739a513eb0de1a2350ab9f741"
)

type ScholarLabs struct {
	service service.Service
	jwt     util.JWTProvider
	cfg     config.Config

	proto.UnimplementedUserServer
}

func New(cfg config.Config, ds store.DataStore) (*ScholarLabs, error) {
	return &ScholarLabs{
		service: service.New(ds),
		jwt:     util.NewJWT(),
		cfg:     cfg,
	}, nil
}

func (p *ScholarLabs) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	credentials, err := p.service.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !hash.ValidatePassword(req.Password, credentials.Password) {
		return nil, fmt.Errorf("unauthorized")
	}

	userToken, err := p.jwt.CreateSignedToken(credentials.UUID,
		time.Now().Add(time.Hour).Unix(),
		issuer,
		[]byte(p.cfg.TokenSecret))
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		Token: userToken,
	}, nil
}

func (p *ScholarLabs) Register(ctx context.Context, req *proto.RegisterRequest) (*emptypb.Empty, error) {
	user, _ := p.service.GetUserByEmail(ctx, req.Email)
	if user != nil {
		return nil, fmt.Errorf("user already exists")
	}

	_, err := p.service.CreateNewUser(ctx, store.User{
		Email: req.Email,

		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
