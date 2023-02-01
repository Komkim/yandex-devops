package app

import (
	"context"
	"log"
	"math/rand"
	"runtime"
	"time"
	"yandex-devops/config"
	"yandex-devops/internal/agent/services"
	myclient "yandex-devops/provider"
)

func Run(ctx *context.Context, config *config.Config) {
	var runtimeStats runtime.MemStats
	var counter int64
	ticker := time.NewTicker(config.Report)

	memStorage := myclient.New(&config.HTTP)

	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			<-ticker.C
			rnd := rand.Float64()
			err := services.Report(memStorage, runtimeStats, counter, rnd)
			if err != nil {
				log.Println(err)
			}
			counter = 0
		}
	}()

	for {
		runtime.ReadMemStats(&runtimeStats)

		counter++

		time.Sleep(config.Poll)
	}

}
