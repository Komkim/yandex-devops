// Создание и работа с сервером
package server

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"yandex-devops/config"
)

// Server - параметры сервера
type Server struct {
	//httpServer - http сервер
	httpServer *http.Server
}

// NewServer - создание нового сервера
func NewServer(cfg *config.HTTP, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			//Addr:    config.Host + ":" + config.Port,
			Addr:    cfg.Address,
			Handler: handler,
		},
	}
}

// Start - запуск сервера
func (s *Server) Start() {
	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

// Stop - остановка сервера
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// GetServer - получение сервера
func (s *Server) GetServer() *http.Server {
	return s.httpServer
}
