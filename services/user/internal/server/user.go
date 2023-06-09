package server

import (
	"context"
	"fmt"
	"time"

	"github.com/FoxFurry/scholarlabs/services/common/hash"
	"github.com/FoxFurry/scholarlabs/services/user/internal/store"
	"github.com/FoxFurry/scholarlabs/services/user/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *ScholarLabs) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	credentials, err := p.service.GetUserByEmail(ctx, req.Email)
	if err != nil {
		p.lg.WithError(err).WithField("user email", req.Email).Error("failed to get user by email")
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
		p.lg.WithError(err).WithField("user email", req.Email).Error("failed to generate JWT token")
		return nil, err
	}

	return &proto.LoginResponse{
		Token: userToken,
	}, nil
}

func (p *ScholarLabs) Register(ctx context.Context, req *proto.RegisterRequest) (*emptypb.Empty, error) {
	user, _ := p.service.GetUserByEmail(ctx, req.Email)
	if user != nil {
		p.lg.WithField("user email", req.Email).Error("user already exists")
		return nil, fmt.Errorf("user already exists")
	}

	_, err := p.service.CreateNewUser(ctx, store.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		p.lg.WithError(err).WithField("user email", req.Email).Error("failed to create new user")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
