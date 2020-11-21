package http

import (
	"context"
)

// Service embed a http server.
type Service struct {
}

// Dial connects http server.
func (s *Service) Dial(ctx context.Context, cfg Config) error {
	return nil
}

func (s *Service) Close(ctx context.Context) error {
	return nil
}
