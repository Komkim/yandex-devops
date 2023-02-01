package agent

import (
	"runtime"
	myclient "yandex-devops/provider"
)

const GAUGE = "gauge"
const COUNTER = "counter"

func ConvertRuntumeStatsToStorageMetrics(stats *runtime.MemStats, counter int64, rand float64) *[]myclient.Metrics {
	metrics := make([]myclient.Metrics, 0, 29)
	var value float64

	value = float64(stats.Alloc)
	metrics = append(metrics, myclient.Metrics{
		ID:    "Alloc",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.BuckHashSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "BuckHashSys",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.Frees)
	metrics = append(metrics, myclient.Metrics{
		ID:    "Frees",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.GCCPUFraction)
	metrics = append(metrics, myclient.Metrics{
		ID:    "GCCPUFraction",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.GCSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "GCSys",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.HeapAlloc)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapAlloc",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.HeapIdle)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapIdle",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.HeapInuse)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapInuse",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.HeapObjects)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapObjects",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.HeapReleased)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapReleased",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.HeapSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapSys",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.LastGC)
	metrics = append(metrics, myclient.Metrics{
		ID:    "LastGC",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.Lookups)
	metrics = append(metrics, myclient.Metrics{
		ID:    "Lookups",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.MCacheInuse)
	metrics = append(metrics, myclient.Metrics{
		ID:    "MCacheInuse",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.MCacheSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "MCacheSys",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.MSpanInuse)
	metrics = append(metrics, myclient.Metrics{
		ID:    "MSpanInuse",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.MSpanSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "MSpanSys",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.Mallocs)
	metrics = append(metrics, myclient.Metrics{
		ID:    "Mallocs",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.NextGC)
	metrics = append(metrics, myclient.Metrics{
		ID:    "NextGC",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.NumForcedGC)
	metrics = append(metrics, myclient.Metrics{
		ID:    "NumForcedGC",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.NumGC)
	metrics = append(metrics, myclient.Metrics{
		ID:    "NumGC",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.OtherSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "OtherSys",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.PauseTotalNs)
	metrics = append(metrics, myclient.Metrics{
		ID:    "PauseTotalNs",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.StackInuse)
	metrics = append(metrics, myclient.Metrics{
		ID:    "StackInuse",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.StackSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "StackSys",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.Sys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "Sys",
		MType: GAUGE,
		Value: &value,
	})

	value = float64(stats.TotalAlloc)
	metrics = append(metrics, myclient.Metrics{
		ID:    "TotalAlloc",
		MType: GAUGE,
		Value: &value,
	})

	metrics = append(metrics, myclient.Metrics{
		ID:    "PollCount",
		MType: COUNTER,
		Delta: &counter,
	})

	metrics = append(metrics, myclient.Metrics{
		ID:    "RandomValue",
		MType: GAUGE,
		Value: &rand,
	})

	return &metrics
}
