package agent

import (
	"context"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
	"math/rand"
	"runtime"
	"time"
	"yandex-devops/config"
	myclient "yandex-devops/provider"
)

type Agent struct {
	cfg        *config.Agent
	updateChan chan []myclient.Metrics
	sendChan   chan myclient.Metrics
}

func NewAgen(cfg *config.Agent, updateChan chan []myclient.Metrics, sendChan chan myclient.Metrics) *Agent {
	return &Agent{
		cfg:        cfg,
		updateChan: updateChan,
		sendChan:   sendChan,
	}
}

func (a *Agent) SendMetric(ctx context.Context, cfg *config.Agent, client *myclient.MyClient) {
	ticker := time.NewTicker(a.cfg.Report)
	var metrics []myclient.Metrics
	go a.sendMetric(cfg.RateLimit, client)

	for {
		select {

		case <-ticker.C:
			for _, m := range metrics {
				a.sendChan <- m
			}

		case metrics = <-a.updateChan:

		case <-ctx.Done():
			return
		}
	}
}

func (a *Agent) sendMetric(limitWorker int, client *myclient.MyClient) {
	for i := 0; i < limitWorker; i++ {
		go func(sChan <-chan myclient.Metrics) {
			for {
				metric, ok := <-sChan
				if ok {
					log.Println("agent send metric")
					log.Println(metric)
					err := client.SendOneMetric(metric)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}(a.sendChan)
	}
}

func (a *Agent) UpdateMetric(ctx context.Context) {
	var runtimeStats runtime.MemStats
	ticker := time.NewTicker(a.cfg.Poll)
	rand.Seed(time.Now().UnixNano())
	var counter int64

	for {
		select {

		case <-ticker.C:
			rnd := rand.Float64()
			counter++
			runtime.ReadMemStats(&runtimeStats)
			a.updateChan <- ConvertRuntumeStatsToStorageMetrics(&runtimeStats, counter, rnd, a.cfg.Key)
		case <-ctx.Done():
			return
		}
	}
}

func (a *Agent) UpdateVirtualMemory(ctx context.Context) {
	ticker := time.NewTicker(a.cfg.Poll)

f:
	for {
		select {

		case <-ticker.C:
			virtualMemory, err := mem.VirtualMemory()
			if err != nil {
				log.Println(err)
				continue f
			}
			a.updateChan <- ConvertVirtualMemoryToStorageMertics(virtualMemory, a.cfg.Key)
		case <-ctx.Done():
			return
		}
	}

}
