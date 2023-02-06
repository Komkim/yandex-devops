package server

import (
	"context"
	"net/http"
	"yandex-devops/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.HTTP, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			//Addr:    config.Host + ":" + config.Port,
			Addr:    cfg.Address,
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

func (s *Server) GetServer() *http.Server {
	return s.httpServer
}
