package services

import (
	transport "Komkim/go-musthave-devops-tpl/cmd/agent/pkg"
	"Komkim/go-musthave-devops-tpl/cmd/agent/storage"
	"runtime"
)

func Report(client transport.MyClient, stats runtime.MemStats, count int, rand float64) {

	m := myStatsConversionFromRuntimeMemStats(stats, count, rand)

	storage.SendAll(client, m.convertToOneMetricSlice())
}
