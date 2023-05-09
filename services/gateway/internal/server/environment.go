package server

import (
	"github.com/FoxFurry/scholarlabs/services/environment/proto"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/httperr"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *ScholarLabs) CreateEnvironment(ctx *gin.Context) {
	var (
		err      error
		env      models.Environment
		userUUID string
	)

	if err = ctx.BindJSON(&env); err != nil {
		s.lg.WithError(err).Error("failed to bind request")
		httperr.Handle(ctx, httperr.New("bad request", 400))
		return
	}

	userUUID, err = s.getUUIDFromContext(ctx)
	if err != nil {
		s.lg.WithError(err).Error("failed to get uuid from request")
		httperr.Handle(ctx, httperr.New("unauthorized", 401))
		return
	}

	if _, err = s.environmentService.CreateEnvironment(ctx, &proto.CreateEnvironmentRequest{
		EnvIdentifier: env.EnvIdentifier,
		Type:          env.Type,
		Name:          env.Name,
		UserUUID:      userUUID,
	}); err != nil {
		s.lg.WithError(err).Error("failed to create environment")
		httperr.Handle(ctx, httperr.New("something went wrong", 500))
		return
	}

	ctx.Status(200)
}

func (s *ScholarLabs) GetEnvironmentsForUser(ctx *gin.Context) {
	userUUID, err := s.getUUIDFromContext(ctx)
	if err != nil {
		s.lg.WithError(err).Error("failed to get uuid from request")
		httperr.Handle(ctx, httperr.New("unauthorized", 401))
		return
	}

	envs, err := s.environmentService.GetEnvironmentsForUser(ctx, &proto.GetEnvironmentsForUserRequest{
		UserUUID: userUUID,
	})
	if err != nil {
		s.lg.WithError(err).Error("failed to get environments for user")
		httperr.Handle(ctx, httperr.New("something went wrong", 500))
		return
	}

	s.lg.Printf("With getter: %d", len(envs.GetEnvironments()))
	s.lg.Printf("Without getter: %d", len(envs.Environments))

	if len(envs.Environments) == 0 {
		ctx.String(404, "no envs found")
		return
	}

	ctx.JSON(200, protoToModelEnvs(envs.Environments))
}

func protoToModelEnv(env *proto.Env) models.Environment {
	var modelEnv models.Environment

	modelEnv.EnvIdentifier = env.GetEnvIdentifier()
	modelEnv.Type = env.GetType()
	modelEnv.Name = env.GetName()

	return modelEnv
}

func protoToModelEnvs(envs []*proto.Env) []models.Environment {
	var modelEnvs = make([]models.Environment, 0, len(envs))

	for _, env := range envs {
		modelEnvs = append(modelEnvs, protoToModelEnv(env))
	}

	return modelEnvs
}
