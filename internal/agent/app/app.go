package app

import (
	"math/rand"
	"runtime"
	"time"
	"yandex-devops/config"
	transport "yandex-devops/internal/agent/pkg"
	"yandex-devops/internal/agent/services"
)

func Run(config *config.Config) {
	var runtimeStats runtime.MemStats
	var counter int

	storage := transport.New(config)

	rand.Seed(time.Now().UnixNano())

	for {
		runtime.ReadMemStats(&runtimeStats)

		counter++

		rnd := rand.Float64()

		if r := counter % config.Report; r == 0 {
			services.Report(storage, runtimeStats, counter, rnd)
			counter = 0
		}

		time.Sleep(time.Second * time.Duration(config.Poll))
	}

}
