package file

import (
	"time"
	"yandex-devops/config"
)

func getTestFileRepo() *FileStorage {
	cfg := config.Server{
		FilePath:     "/tmp/devops-metrics-db.json",
		FileInterval: 300 * time.Second,
		FileRestore:  true,
		Key:          "123",
	}

	return NewFileStorage(&cfg)
}
