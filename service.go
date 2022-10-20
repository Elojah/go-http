package http

import (
	"context"
	"crypto/tls"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/acme/autocert"
)

// Service embed a http server.
type Service struct {
	*http.Server
	*mux.Router
}

// Dial connects http server.
func (s *Service) Dial(ctx context.Context, cfg Config) error {
	s.Router = mux.NewRouter()

	var tlscfg *tls.Config

	if len(cfg.HostWhitelist) > 0 {
		c := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(cfg.HostWhitelist...),
		}

		tlscfg = &tls.Config{GetCertificate: c.GetCertificate, MinVersion: tls.VersionTLS12}
	} else {
		cert, err := tls.LoadX509KeyPair(cfg.Cert, cfg.Key)
		if err != nil {
			return err
		}

		tlscfg = &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12}
	}

	s.Server = &http.Server{
		Addr:              strings.Join([]string{cfg.Hostname, cfg.Port}, ":"),
		ReadHeaderTimeout: time.Minute,
		TLSConfig:         tlscfg,
		Handler:           s.Router,
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
