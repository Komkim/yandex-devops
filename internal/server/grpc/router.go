package mygrpc

import (
	"yandex-devops/config"
	"yandex-devops/internal/server/service"
	pb "yandex-devops/proto"
)

type Router struct {
	//services - внутренние сервисы сервера для работы с внешними точками взаимоействия сервера
	services *service.Services
	//cfg - конфиги сервера
	cfg *config.Server

	*pb.UnimplementedMetricsServer
}

// NewRouter - создание нового матшрутизатора
func NewRouter(cfg *config.Server, s *service.Services) *Router {
	return &Router{cfg: cfg, services: s}
}
