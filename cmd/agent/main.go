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
		log.Println(err)
		panic(err)
	}

	ch := make(chan *[]myclient.Metrics)
	client := myclient.New(&cfg.HTTP)

	a := agent.NewAgen(ctx, &cfg.Agent, ch)

	go a.UpdateMetric()
	go a.SendMetric(&client)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
