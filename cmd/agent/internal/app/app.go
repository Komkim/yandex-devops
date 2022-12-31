package app

import (
	"Komkim/go-musthave-devops-tpl/cmd/agent/config"
	"Komkim/go-musthave-devops-tpl/cmd/agent/internal/services"
	"Komkim/go-musthave-devops-tpl/cmd/agent/storage"
	"math/rand"
	"runtime"
	"time"
)

func Run(storage storage.Sending) {
	var runtimeStats runtime.MemStats
	var counter int

	rand.Seed(time.Now().UnixNano())

	for {
		runtime.ReadMemStats(&runtimeStats)

		counter++

		rnd := rand.Float64()

		if r := counter % config.ReportInterval; r == 0 {
			services.Report(storage, runtimeStats, counter, rnd)
			counter = 0
		}

		time.Sleep(time.Second * config.PollInterval)
	}

}
