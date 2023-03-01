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

	updateChan := make(chan []myclient.Metrics)
	sendChan := make(chan myclient.Metrics)
	client := myclient.New(&cfg.HTTP)

	a := agent.NewAgen(&cfg.Agent, updateChan, sendChan)

	//go a.UpdateVirtualMemory(ctx)
	go a.UpdateMetric(ctx)

	go a.SendMetric(ctx, &cfg.Agent, &client)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
