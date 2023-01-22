package services

import (
	"runtime"
	"yandex-devops/storage"
)

func Report(storage storage.Storage, stats runtime.MemStats, count int64, rand float64) error {

	m := myStatsConversionFromRuntimeMemStats(stats, count, rand)

	_, err := storage.SetAll(m.convertToOneMetricSlice())
	if err != nil {
		return err
	}
	return nil
}
