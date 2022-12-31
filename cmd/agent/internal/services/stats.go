package services

import (
	"Komkim/go-musthave-devops-tpl/cmd/agent/storage"
	"fmt"
	"reflect"
	"runtime"
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

func (m MyStats) convertToOneMetricSlice() []storage.OneMetric {
	val := reflect.ValueOf(m)
	metrics := make([]storage.OneMetric, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		metrics = append(metrics, storage.OneMetric{
			val.Type().Field(i).Type.String(),
			string(val.Type().Field(i).Name),
			fmt.Sprint(val.Field(i)),
		})
	}
	return metrics
}

func myStatsConversionFromRuntimeMemStats(stats runtime.MemStats, c int, rand float64) MyStats {
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
