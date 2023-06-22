// Agent - модуль для сбора и отправки метрик на серве
package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

		<-quit
		cancel()
		return nil
	})

	cfg, err := config.InitFlagAgent()
	if err != nil {
		log.Panic(err)
	}

	updateRuntimeChan := make(chan []myclient.Metrics)
	updateVirtMemoryChan := make(chan []myclient.Metrics)
	sendChan := make(chan myclient.Metrics)

	a := agent.NewAgent(cfg, updateRuntimeChan, updateVirtMemoryChan, sendChan)

	g.Go(func() error {
		return a.UpdateVirtualMemory(gCtx)
	})

	g.Go(func() error {
		return a.UpdateMetric(gCtx)
	})

	switch cfg.APIType {
	case config.GRPC:
		client := myclient.NewGrpcClient(cfg)
		g.Go(func() error {
			return a.SendMetric(gCtx, cfg, &client)
		})
	default:
		client := myclient.New(cfg)
		g.Go(func() error {
			return a.SendMetric(gCtx, cfg, &client)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
	fmt.Println("Agent done")
}
