package server

import (
	"google.golang.org/grpc"
	"net"
	"yandex-devops/config"
	pb "yandex-devops/proto"
)

type MyGrpcServer struct {
	Handler pb.MetricsServer
	Server  *grpc.Server
	cfg     *config.Server
}

func NewGrpcServer(cfg *config.Server, srv pb.MetricsServer) *MyGrpcServer {
	// создаём gRPC-сервер без зарегистрированной службы
	serv := grpc.NewServer()
	s := &MyGrpcServer{Handler: srv, Server: serv, cfg: cfg}
	// регистрируем сервис
	pb.RegisterMetricsServer(s.Server, s.Handler)

	return s
}

func (gs *MyGrpcServer) Start() error {
	listen, err := net.Listen("tcp", gs.cfg.GrpcAddress)
	if err != nil {
		return err
	}

	err = gs.Server.Serve(listen)
	if err != nil {
		return err
	}
	return nil
}
