package services

import (
	"Komkim/go-musthave-devops-tpl/cmd/agent/storage"
	"runtime"
)

type Reporting interface {
	Report(storage storage.Sending, stats runtime.MemStats, count int, rand float64)
}
