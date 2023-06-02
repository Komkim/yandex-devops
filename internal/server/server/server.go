// Создание и работа с сервером
package server

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	_ "net/http/pprof"
	"yandex-devops/config"
)

// Server - параметры сервера
type Server struct {
	//httpServer - http сервер
	httpServer *http.Server
	//cfg - настройки подключения
	cfg *config.Server
}

// NewServer - создание нового сервера
func NewServer(cfg *config.Server, handler http.Handler) *Server {
	cer, err := tls.LoadX509KeyPair("certificat/local.crt", cfg.CryptoKey)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}

	return &Server{
		httpServer: &http.Server{
			Addr:      cfg.Address,
			Handler:   handler,
			TLSConfig: tlsConfig,
		},
		cfg: cfg,
	}
}

// Start - запуск сервера
func (s *Server) Start() error {
	return s.httpServer.ListenAndServeTLS("certificat/local.crt", s.cfg.CryptoKey)
}

// Stop - остановка сервера
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// GetServer - получение сервера
func (s *Server) GetServer() *http.Server {
	return s.httpServer
}
