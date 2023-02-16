package agent

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"runtime"
	myclient "yandex-devops/provider"
)

const GAUGE = "gauge"
const COUNTER = "counter"

func convert(stats *runtime.MemStats, counter int64, rand float64) []myclient.Metrics {
	metrics := make([]myclient.Metrics, 0, 29)

	Alloc := float64(stats.Alloc)
	metrics = append(metrics, myclient.Metrics{
		ID:    "Alloc",
		MType: GAUGE,
		Value: &Alloc,
	})

	BuckHashSys := float64(stats.BuckHashSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "BuckHashSys",
		MType: GAUGE,
		Value: &BuckHashSys,
	})

	Frees := float64(stats.Frees)
	metrics = append(metrics, myclient.Metrics{
		ID:    "Frees",
		MType: GAUGE,
		Value: &Frees,
	})

	GCCPUFraction := float64(stats.GCCPUFraction)
	metrics = append(metrics, myclient.Metrics{
		ID:    "GCCPUFraction",
		MType: GAUGE,
		Value: &GCCPUFraction,
	})

	GCSys := float64(stats.GCSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "GCSys",
		MType: GAUGE,
		Value: &GCSys,
	})

	HeapAlloc := float64(stats.HeapAlloc)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapAlloc",
		MType: GAUGE,
		Value: &HeapAlloc,
	})

	HeapIdle := float64(stats.HeapIdle)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapIdle",
		MType: GAUGE,
		Value: &HeapIdle,
	})

	HeapInuse := float64(stats.HeapInuse)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapInuse",
		MType: GAUGE,
		Value: &HeapInuse,
	})

	HeapObjects := float64(stats.HeapObjects)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapObjects",
		MType: GAUGE,
		Value: &HeapObjects,
	})

	HeapReleased := float64(stats.HeapReleased)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapReleased",
		MType: GAUGE,
		Value: &HeapReleased,
	})

	HeapSys := float64(stats.HeapSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "HeapSys",
		MType: GAUGE,
		Value: &HeapSys,
	})

	LastGC := float64(stats.LastGC)
	metrics = append(metrics, myclient.Metrics{
		ID:    "LastGC",
		MType: GAUGE,
		Value: &LastGC,
	})

	Lookups := float64(stats.Lookups)
	metrics = append(metrics, myclient.Metrics{
		ID:    "Lookups",
		MType: GAUGE,
		Value: &Lookups,
	})

	MCacheInuse := float64(stats.MCacheInuse)
	metrics = append(metrics, myclient.Metrics{
		ID:    "MCacheInuse",
		MType: GAUGE,
		Value: &MCacheInuse,
	})

	MCacheSys := float64(stats.MCacheSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "MCacheSys",
		MType: GAUGE,
		Value: &MCacheSys,
	})

	MSpanInuse := float64(stats.MSpanInuse)
	metrics = append(metrics, myclient.Metrics{
		ID:    "MSpanInuse",
		MType: GAUGE,
		Value: &MSpanInuse,
	})

	MSpanSys := float64(stats.MSpanSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "MSpanSys",
		MType: GAUGE,
		Value: &MSpanSys,
	})

	Mallocs := float64(stats.Mallocs)
	metrics = append(metrics, myclient.Metrics{
		ID:    "Mallocs",
		MType: GAUGE,
		Value: &Mallocs,
	})

	NextGC := float64(stats.NextGC)
	metrics = append(metrics, myclient.Metrics{
		ID:    "NextGC",
		MType: GAUGE,
		Value: &NextGC,
	})

	NumForcedGC := float64(stats.NumForcedGC)
	metrics = append(metrics, myclient.Metrics{
		ID:    "NumForcedGC",
		MType: GAUGE,
		Value: &NumForcedGC,
	})

	NumGC := float64(stats.NumGC)
	metrics = append(metrics, myclient.Metrics{
		ID:    "NumGC",
		MType: GAUGE,
		Value: &NumGC,
	})

	OtherSys := float64(stats.OtherSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "OtherSys",
		MType: GAUGE,
		Value: &OtherSys,
	})

	PauseTotalNs := float64(stats.PauseTotalNs)
	metrics = append(metrics, myclient.Metrics{
		ID:    "PauseTotalNs",
		MType: GAUGE,
		Value: &PauseTotalNs,
	})

	StackInuse := float64(stats.StackInuse)
	metrics = append(metrics, myclient.Metrics{
		ID:    "StackInuse",
		MType: GAUGE,
		Value: &StackInuse,
	})

	StackSys := float64(stats.StackSys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "StackSys",
		MType: GAUGE,
		Value: &StackSys,
	})

	Sys := float64(stats.Sys)
	metrics = append(metrics, myclient.Metrics{
		ID:    "Sys",
		MType: GAUGE,
		Value: &Sys,
	})

	TotalAlloc := float64(stats.TotalAlloc)
	metrics = append(metrics, myclient.Metrics{
		ID:    "TotalAlloc",
		MType: GAUGE,
		Value: &TotalAlloc,
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

	return metrics
}

func generateHas(key string, m []myclient.Metrics) []myclient.Metrics {
	if len(key) <= 0 {
		return m
	}

	for k, v := range m {
		var data []byte
		switch m[k].MType {
		case COUNTER:
			data = []byte(fmt.Sprintf("%s:%s:%d", v.ID, v.MType, *v.Delta))
		case GAUGE:
			data = []byte(fmt.Sprintf("%s:%s:%f", v.ID, v.MType, *v.Value))
		}

		h := hmac.New(sha256.New, []byte(key))
		h.Write(data)
		dst := h.Sum(nil)

		m[k].Hash = hex.EncodeToString(dst)
	}

	return m
}

func ConvertRuntumeStatsToStorageMetrics(stats *runtime.MemStats, counter int64, rand float64, key string) []myclient.Metrics {
	return generateHas(key, convert(stats, counter, rand))
}
