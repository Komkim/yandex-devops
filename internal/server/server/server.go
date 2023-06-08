// Создание и работа с сервером
package server

import (
	"context"
	"crypto/tls"
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

const certFile = "certificat/certificate.crt"

// NewServer - создание нового сервера
func NewServer(cfg *config.Server, handler http.Handler) *Server {
	if len(cfg.CryptoKey) > 0 {

		cer, err := tls.LoadX509KeyPair(certFile, cfg.CryptoKey)
		if err != nil {
			//log.Println(err)
			panic(err)
		}

		tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}

		return &Server{
			httpServer: &http.Server{
				Addr:      cfg.Address,
				Handler:   handler,
				TLSConfig: tlsConfig,
				//TLSConfig: &tls.Config{ServerName: cfg.Address},
			},
			cfg: cfg,
		}
	}
	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.Address,
			Handler: handler,
		},
		cfg: cfg,
	}
}

// Start - запуск сервера
func (s *Server) Start() error {
	if len(s.cfg.CryptoKey) > 0 {
		return s.httpServer.ListenAndServeTLS(certFile, s.cfg.CryptoKey)
	}
	return s.httpServer.ListenAndServe()
}

// Stop - остановка сервера
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// GetServer - получение сервера
func (s *Server) GetServer() *http.Server {
	return s.httpServer
}
