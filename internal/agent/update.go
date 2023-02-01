package agent

import (
	"context"
	"math/rand"
	"runtime"
	"time"
	"yandex-devops/config"
	myclient "yandex-devops/provider"
)

func UpdateMetric(ctx context.Context, cfg *config.Agent, chanSendMetric chan<- *[]myclient.Metrics, counter *Counter) {
	var runtimeStats runtime.MemStats
	ticker := time.NewTicker(cfg.Poll)
	rand.Seed(time.Now().UnixNano())

	for {
		select {
		case <-ticker.C:
			rnd := rand.Float64()
			counter.Inc()
			runtime.ReadMemStats(&runtimeStats)
			chanSendMetric <- ConvertRuntumeStatsToStorageMetrics(&runtimeStats, counter.Get(), rnd)
		case <-ctx.Done():
			return
		}
	}
}
