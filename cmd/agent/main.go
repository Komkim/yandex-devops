package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"yandex-devops/config"
	"yandex-devops/internal/agent"
	myclient "yandex-devops/provider"
)

func main() {
	ctx, cencel := context.WithCancel(context.Background())
	defer cencel()

	cfg, err := config.InitAgent()
	//cfg, err := config.InitFlagAgent()
	if err != nil {
		log.Println(err)
	}

	counter := agent.NewCounter()
	ch := make(chan *[]myclient.Metrics)
	client := myclient.New(&cfg.HTTP)

	go agent.UpdateMetric(ctx, &cfg.Agent, ch, counter)
	go func(ctx context.Context) {
		err := agent.SendMetric(ctx, &cfg.Agent, &client, ch, counter)
		if err != nil {
			log.Println(err)
		}
	}(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
