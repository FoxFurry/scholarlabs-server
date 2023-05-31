package service

import (
	"context"
	"fmt"
	"log"

	"github.com/FoxFurry/scholarlabs/virt"
)

func (s *service) BidirectionalTerminal(ctx context.Context, engine, termRef string, execCmd []string) (virt.Terminal, error) {
	targetEngine, err := s.resolveEngine(engine)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve engine: %w", err)
	}

	log.Println("ExecCmd", execCmd)

	terminal, err := targetEngine.StartTerminal(ctx, termRef, execCmd)
	if err != nil {
		return nil, fmt.Errorf("failed to start terminal: %w", err)
	}

	return terminal, nil
}
