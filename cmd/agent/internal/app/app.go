package app

import (
	"Komkim/go-musthave-devops-tpl/cmd/agent/config"
	"Komkim/go-musthave-devops-tpl/cmd/agent/internal/services"
	transport "Komkim/go-musthave-devops-tpl/cmd/agent/pkg"
	"math/rand"
	"runtime"
	"time"
)

func Run(client transport.MyClient) {
	var runtimeStats runtime.MemStats
	var counter int

	rand.Seed(time.Now().UnixNano())

m:
	for {
		runtime.ReadMemStats(&runtimeStats)

		counter++

		rnd := rand.Float64()

		if r := counter % config.ReportInterval; r == 0 {
			services.Report(client, runtimeStats, counter, rnd)
		}

		time.Sleep(time.Second * config.PollInterval)

		if counter == 30 {
			break m
		}
	}

}
