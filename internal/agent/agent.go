// Модуль для сбора и отправки метрик
package agent

import (
	"context"
	"log"
	"math/rand"
	"runtime"
	"time"
	"yandex-devops/config"
	myclient "yandex-devops/provider"

	"github.com/shirou/gopsutil/v3/mem"
)

// Agent - агент для сбора и отправки метрик
type Agent struct {
	//cfg - параметры агенты
	cfg *config.Agent
	//updateRuntimeChan - канал для обновления основных метрик
	updateRuntimeChan chan []myclient.Metrics
	//updateVirtMemoryChan - канал для обновления метрик памяти
	updateVirtMemoryChan chan []myclient.Metrics
	//sendChan - канал для отправки метрик
	sendChan chan myclient.Metrics
}

// NewAgent - создание нового агента
func NewAgent(cfg *config.Agent, updateRuntimeChan chan []myclient.Metrics, updateVirtMemoryChan chan []myclient.Metrics, sendChan chan myclient.Metrics) *Agent {
	return &Agent{
		cfg:                  cfg,
		updateRuntimeChan:    updateRuntimeChan,
		updateVirtMemoryChan: updateVirtMemoryChan,
		sendChan:             sendChan,
	}
}

// SendMetric - отправка метрик на сервер
func (a *Agent) SendMetric(ctx context.Context, cfg *config.Agent, client *myclient.MyClient) {
	ticker := time.NewTicker(a.cfg.Report)
	var metricsRuntime []myclient.Metrics
	var metricsVirtMemory []myclient.Metrics
	go a.sendMetric(cfg.RateLimit, client)

	for {
		select {

		case <-ticker.C:
			for _, m := range metricsRuntime {
				a.sendChan <- m
			}
			for _, m := range metricsVirtMemory {
				a.sendChan <- m
			}

		case metricsRuntime = <-a.updateRuntimeChan:
		case metricsVirtMemory = <-a.updateVirtMemoryChan:

		case <-ctx.Done():
			return
		}
	}
}

// sendMetric - запуск необходимого колличества воркеров для отправки метрик на сервер
func (a *Agent) sendMetric(limitWorker int, client *myclient.MyClient) {
	for i := 0; i < limitWorker; i++ {
		go func(sChan <-chan myclient.Metrics) {
			for {
				metric, ok := <-sChan
				if ok {
					err := client.SendOneMetric(metric)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}(a.sendChan)
	}
}

// UpdateMetric - обновление основных метрик
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
			a.updateRuntimeChan <- ConvertRuntumeStatsToStorageMetrics(&runtimeStats, counter, rnd, a.cfg.Key)
		case <-ctx.Done():
			return
		}
	}
}

// UpdateVirtualMemory - обновление метрик памяти
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
			a.updateVirtMemoryChan <- ConvertVirtualMemoryToStorageMertics(virtualMemory, a.cfg.Key)
		case <-ctx.Done():
			return
		}
	}

}
