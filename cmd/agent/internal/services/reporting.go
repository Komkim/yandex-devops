package services

import (
	"cmd/storage"
	"runtime"
)

type Reporting interface {
	Report(storage storage.Sending, stats runtime.MemStats, count int, rand float64)
}
