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

	cfg, err := config.InitFlagAgent()
	if err != nil {
		log.Panic(err)
	}

	ch := make(chan []myclient.Metrics, 10)
	client := myclient.New(&cfg.HTTP)

	a := agent.NewAgen(&cfg.Agent, ch)

	go a.UpdateMetric(ctx)
	go a.UpdateVirtualMemory(ctx)
	for i := 0; i < cfg.Agent.RateLimit; i++ {
		go a.SendMetric(ctx, &client)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
