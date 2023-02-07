package agent

import (
	"context"
	"log"
	"math/rand"
	"runtime"
	"time"
	"yandex-devops/config"
	myclient "yandex-devops/provider"
)

type Agent struct {
	ctx context.Context
	cfg *config.Agent
	sm  chan *[]myclient.Metrics
}

func NewAgen(ctx context.Context, cfg *config.Agent, ch chan *[]myclient.Metrics) *Agent {
	return &Agent{
		ctx: ctx,
		cfg: cfg,
		sm:  ch,
	}
}

func (a *Agent) SendMetric(client *myclient.MyClient) {
	ticker := time.NewTicker(a.cfg.Report)
	var metrics *[]myclient.Metrics

	for {
		select {

		case <-ticker.C:
			err := client.SendAllMetric(metrics)
			if err != nil {
				log.Println(err)
			}

		case metrics = <-a.sm:

		case <-a.ctx.Done():
			return
		}
	}
}

func (a *Agent) UpdateMetric() {
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
		case <-a.ctx.Done():
			return
		}
	}
}
