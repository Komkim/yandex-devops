package service

import (
	"yandex-devops/config"
	"yandex-devops/storage"
)

const GAUGE = "gauge"
const COUNTER = "counter"

type Services struct {
	Mss            *MemStorageService
	StorageService *StorageService
}

func NewServices(cfg *config.Server, memory, storage storage.Storage) *Services {
	memoryService := NewMemStorageService(&memory)
	return &Services{
		Mss:            memoryService,
		StorageService: NewStorageService(cfg, &storage, memoryService),
	}
}
