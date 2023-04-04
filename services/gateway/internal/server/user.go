package server

import (
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/httperr"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/models"
	"github.com/FoxFurry/scholarlabs/services/user/proto"
	"github.com/gin-gonic/gin"
)

func (s *ScholarLabs) Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		s.lg.WithError(err).WithField("user email", user.Email).Error("failed to bind request")
		httperr.Handle(ctx, httperr.New("bad request", 400))
		return
	}

	if _, err := s.userService.Register(ctx, &proto.RegisterRequest{
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}); err != nil {
		s.lg.WithError(err).WithField("user email", user.Email).Error("failed to register new user")
		httperr.Handle(ctx, httperr.New("oops, something went wrong", 500))
		return
	}

	ctx.String(200, "successfully created new user")
}

func (s *ScholarLabs) Login(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		s.lg.WithError(err).WithField("user email", user.Email).Error("failed to bind request")
		httperr.Handle(ctx, httperr.New("bad request", 400))
		return
	}

	loginResponse, err := s.userService.Login(ctx, &proto.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		s.lg.WithError(err).WithField("user email", user.Email).Error("failed to login")
		httperr.Handle(ctx, httperr.New("oops, something went wrong", 500))
		return
	}

	ctx.JSON(200, gin.H{
		"token": loginResponse.Token,
	})
}
