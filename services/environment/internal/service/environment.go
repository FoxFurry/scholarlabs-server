package service

import (
	"context"
	"fmt"

	"github.com/FoxFurry/scholarlabs/services/environment/internal/store"
	"github.com/FoxFurry/scholarlabs/virt"
	"github.com/google/uuid"
)

func (s *service) CreateEnvironment(ctx context.Context, c store.Environment) (*store.Environment, error) {
	prototype, err := s.db.GetPrototypeByUUID(ctx, c.PrototypeUUID)
	if err != nil {
		return nil, handleDBError(err, "could not create environment")
	}

	targetEngine, err := s.resolveEngine(prototype.Engine)
	if err != nil {
		return nil, err
	}

	machineUUID, err := targetEngine.Spin(ctx, prototype.EngineRef, prototype.EngineRef)
	if err != nil {
		return nil, fmt.Errorf("could not spin environment: %w", err)
	}

	c.MachineUUID = machineUUID
	c.UUID = uuid.New().String()

	if err := s.db.CreateEnvironment(ctx, c); err != nil {
		return nil, handleDBError(err, "could not create environment")
	}

	return &c, nil
}

func (s *service) GetEnvironmentsForUser(ctx context.Context, userUUID string) ([]store.Environment, error) {
	courses, err := s.db.GetEnvironmentsForUser(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (s *service) GetEnvironmentByUUID(ctx context.Context, envUUID string) (*store.Environment, error) {
	env, err := s.db.GetEnvironmentDetails(ctx, envUUID)
	if err != nil {
		return nil, err
	}

	return env, nil
}

func (s *service) resolveEngine(engine string) (virt.Engine, error) {
	switch engine {
	case "CONTAINER":
		return s.dockerEngine, nil
	case "VIRTUAL_MACHINE":
		return nil, fmt.Errorf("virtual machines are not yet implemented")
	default:
		return nil, fmt.Errorf("bad engine type")
	}
}
