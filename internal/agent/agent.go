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
	cfg *config.Agent
	sm  chan []myclient.Metrics
}

func NewAgen(cfg *config.Agent, ch chan []myclient.Metrics) *Agent {
	return &Agent{
		cfg: cfg,
		sm:  ch,
	}
}

func (a *Agent) SendMetric(ctx context.Context, client *myclient.MyClient) {
	ticker := time.NewTicker(a.cfg.Report)
	var metrics []myclient.Metrics

	for {
		select {

		case <-ticker.C:
			err := client.SendAllMetric(metrics)
			if err != nil {
				log.Println(err)
			}

		case metrics = <-a.sm:

		case <-ctx.Done():
			return
		}
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
			a.sm <- ConvertRuntumeStatsToStorageMetrics(&runtimeStats, counter, rnd, a.cfg.Key)
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
			a.sm <- ConvertVirtualMemoryToStorageMertics(virtualMemory, a.cfg.Key)
		case <-ctx.Done():
			return
		}
	}

}
