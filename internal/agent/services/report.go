package services

import (
	"runtime"
	"yandex-devops/internal/agent/storage"
)

func Report(storage storage.Sending, stats runtime.MemStats, count int, rand float64) {

	m := myStatsConversionFromRuntimeMemStats(stats, count, rand)

	storage.SendAll(m.convertToOneMetricSlice())
}
