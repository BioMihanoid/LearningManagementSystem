package pkg

import (
	"context"
	"github.com/BioMihanoid/LearningManagementSystem/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start(cfg config.Config, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + cfg.PortServ,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
		Handler:      handler,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
