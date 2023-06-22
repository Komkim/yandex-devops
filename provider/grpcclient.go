package myclient

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	"yandex-devops/config"
	pb "yandex-devops/proto"
)

type GrpcClient struct {
	cfg    *config.Agent
	client pb.MetricsClient
}

func NewGrpcClient(cfg *config.Agent) GrpcClient {
	conn, err := grpc.Dial(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return GrpcClient{
		cfg:    cfg,
		client: pb.NewMetricsClient(conn),
	}
}

// SendOneMetric - отправка одной метрики
func (c GrpcClient) SendOneMetric(metric Metrics) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	m := &pb.SaveOrUpdateRequest{
		Id:   metric.ID,
		Type: metric.MType,
		Hash: metric.Hash,
	}
	if metric.Value != nil {
		m.Value = &wrappers.DoubleValue{Value: *metric.Value}
	}
	if metric.Delta != nil {
		m.Delta = &wrappers.Int64Value{Value: *metric.Delta}
	}

	c.client.SaveOrUpdate(ctx, m)
	return nil
}

// SendAllMetric - отправка нескольких метрик
func (c GrpcClient) SendAllMetric(metrics []Metrics) error {
	return nil
}
