package services

import (
	"runtime"
	"yandex-devops/internal/agent/storage"
)

type Reporting interface {
	Report(storage storage.Sending, stats runtime.MemStats, count int, rand float64)
}
