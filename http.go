package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	if len(cfg.AllowedCORS) > 0 {
		c := cors.New(cors.Options{
			AllowedOrigins: cfg.AllowedCORS,
		})

		s.Server.Handler = c.Handler(s.Server.Handler)
	}

	return nil
}

func (s *Service) Close(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
