package services

import (
	"runtime"
	"yandex-devops/storage"
)

type Reporting interface {
	Report(storage storage.Storage, stats runtime.MemStats, count int, rand float64)
}
