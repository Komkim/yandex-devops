package server

import (
	"context"
	"log"
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

func (s *Server) Start() {
	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) GetServer() *http.Server {
	return s.httpServer
}
