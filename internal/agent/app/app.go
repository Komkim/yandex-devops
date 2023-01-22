package app

import (
	"log"
	"math/rand"
	"runtime"
	"time"
	"yandex-devops/config"
	transport "yandex-devops/internal/agent/pkg"
	"yandex-devops/internal/agent/services"
)

func Run(config *config.Config) {
	var runtimeStats runtime.MemStats
	var counter int64
	ticker := time.NewTicker(config.Report)

	storage := transport.New(config)

	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			<-ticker.C
			rnd := rand.Float64()
			err := services.Report(storage, runtimeStats, counter, rnd)
			if err != nil {
				log.Println(err)
			}
			counter = 0
		}
	}()

	for {
		runtime.ReadMemStats(&runtimeStats)

		counter++
		//
		//rnd := rand.Float64()
		//
		//
		//
		//if r := counter % config.Report; r == 0 {
		//	err := services.Report(storage, runtimeStats, counter, rnd)
		//	if err != nil {
		//		log.Println(err)
		//	}
		//	counter = 0
		//}

		time.Sleep(config.Poll)
	}

}
