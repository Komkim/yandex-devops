package services

import (
	"reflect"
	"runtime"
	"strings"
	"yandex-devops/storage"
)

type gauge float64

func (g *gauge) string() string {
	return "gauge"
}

type counter int

func (c *counter) string() string {
	return "counter"
}

type MyStats struct {
	Alloc         gauge
	BuckHashSys   gauge
	Frees         gauge
	GCCPUFraction gauge
	GCSys         gauge
	HeapAlloc     gauge
	HeapIdle      gauge
	HeapInuse     gauge
	HeapObjects   gauge
	HeapReleased  gauge
	HeapSys       gauge
	LastGC        gauge
	Lookups       gauge
	MCacheInuse   gauge
	MCacheSys     gauge
	MSpanInuse    gauge
	MSpanSys      gauge
	Mallocs       gauge
	NextGC        gauge
	NumForcedGC   gauge
	NumGC         gauge
	OtherSys      gauge
	PauseTotalNs  gauge
	StackInuse    gauge
	StackSys      gauge
	Sys           gauge
	TotalAlloc    gauge

	PollCount   counter
	RandomValue gauge
}

func (m MyStats) convertToOneMetricSlice() []storage.Metrics {
	val := reflect.ValueOf(m)
	metrics := make([]storage.Metrics, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		id := val.Type().Field(i).Name
		mtype := strings.Replace(val.Type().Field(i).Type.String(), "services.", "", -1)
		switch mtype {
		case "gauge":
			value := val.Field(i).Float()
			metrics = append(metrics, storage.Metrics{
				ID:    id,
				MType: mtype,
				Value: &value,
			})
		case "counter":
			delta := val.Field(i).Int()
			metrics = append(metrics, storage.Metrics{
				ID:    id,
				MType: mtype,
				Delta: &delta,
			})
		}
	}
	return metrics

}

func myStatsConversionFromRuntimeMemStats(stats runtime.MemStats, c int64, rand float64) MyStats {
	return MyStats{
		Alloc:         gauge(stats.Alloc),
		BuckHashSys:   gauge(stats.BuckHashSys),
		Frees:         gauge(stats.Frees),
		GCCPUFraction: gauge(stats.GCCPUFraction),
		GCSys:         gauge(stats.GCSys),
		HeapAlloc:     gauge(stats.HeapAlloc),
		HeapIdle:      gauge(stats.HeapIdle),
		HeapInuse:     gauge(stats.HeapInuse),
		HeapObjects:   gauge(stats.HeapObjects),
		HeapReleased:  gauge(stats.HeapReleased),
		HeapSys:       gauge(stats.HeapSys),
		LastGC:        gauge(stats.LastGC),
		Lookups:       gauge(stats.Lookups),
		MCacheInuse:   gauge(stats.MCacheInuse),
		MCacheSys:     gauge(stats.MCacheSys),
		MSpanInuse:    gauge(stats.MSpanInuse),
		MSpanSys:      gauge(stats.MSpanSys),
		Mallocs:       gauge(stats.Mallocs),
		NextGC:        gauge(stats.NextGC),
		NumForcedGC:   gauge(stats.NumForcedGC),
		NumGC:         gauge(stats.NumGC),
		OtherSys:      gauge(stats.OtherSys),
		PauseTotalNs:  gauge(stats.PauseTotalNs),
		StackInuse:    gauge(stats.StackInuse),
		StackSys:      gauge(stats.StackSys),
		Sys:           gauge(stats.Sys),
		TotalAlloc:    gauge(stats.TotalAlloc),

		PollCount:   counter(c),
		RandomValue: gauge(rand),
	}
}

func convert(stats runtime.MemStats, count int64, rand float64) *[]storage.Metrics {

	val := reflect.ValueOf(stats)
	metrics := make([]storage.Metrics, 0, val.NumField()+2)

	for i := 0; i < val.NumField(); i++ {
		value := val.Field(i).Interface().(float64)
		metrics = append(metrics, storage.Metrics{
			ID:    string(val.Type().Field(i).Name),
			MType: "gauge",
			Value: &value,
		})
	}

	metrics = append(metrics, storage.Metrics{
		ID:    "PollCount",
		MType: "counter",
		Delta: &count,
	})

	metrics = append(metrics, storage.Metrics{
		ID:    "RandomValue",
		MType: "gauge",
		Value: &rand,
	})

	return &metrics
}
