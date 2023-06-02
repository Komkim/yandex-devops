// Agent - модуль для сбора и отправки метрик на серве
package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"yandex-devops/config"
	"yandex-devops/internal/agent"
	myclient "yandex-devops/provider"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	fmt.Printf("Build version: %s", buildVersion)
	fmt.Println()
	fmt.Printf("Build date: %s", buildDate)
	fmt.Println()
	fmt.Printf("Build commit: %s", buildCommit)
	fmt.Println()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	var wg sync.WaitGroup

	cfg, err := config.InitFlagAgent()
	if err != nil {
		log.Panic(err)
	}

	updateRuntimeChan := make(chan []myclient.Metrics)
	updateVirtMemoryChan := make(chan []myclient.Metrics)
	sendChan := make(chan myclient.Metrics)
	client := myclient.New(cfg)

	a := agent.NewAgent(cfg, updateRuntimeChan, updateVirtMemoryChan, sendChan)

	wg.Add(1)
	go a.UpdateVirtualMemory(ctx, &wg)
	wg.Add(1)
	go a.UpdateMetric(ctx, &wg)
	wg.Add(1)
	go a.SendMetric(ctx, &wg, cfg, &client)

	wg.Wait()
	fmt.Println("Agent done")
}
