package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Service embed a http server.
type Service struct {
	*http.Server
	*mux.Router
}

// Dial connects http server.
func (s *Service) Dial(ctx context.Context, cfg Config) error {
	s.Router = mux.NewRouter()

	s.Server = &http.Server{
		Addr:    strings.Join([]string{cfg.Hostname, cfg.Port}, ":"),
		Handler: s.Router,
	}

	return nil
}

func (s *Service) Close(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
