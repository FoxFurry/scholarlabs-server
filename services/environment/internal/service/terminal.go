package service

import (
	"context"
	"fmt"

	"github.com/FoxFurry/scholarlabs/virt"
)

func (s *service) BidirectionalTerminal(ctx context.Context, engine, termRef string) (virt.Terminal, error) {
	targetEngine, err := s.resolveEngine(engine)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve engine: %w", err)
	}

	terminal, err := targetEngine.StartTerminal(ctx, termRef)
	if err != nil {
		return nil, fmt.Errorf("failed to start terminal: %w", err)
	}

	return terminal, nil
}
