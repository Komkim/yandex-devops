package server

import (
	"context"
	"net/http"
	"yandex-devops/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(config *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			//Addr:    config.Host + ":" + config.Port,
			Addr:    config.Address,
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
