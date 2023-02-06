package service

import "yandex-devops/storage"

const GAUGE = "gauge"
const COUNTER = "counter"

type Services struct {
	Mss *MemStorageService
	Fss *FileStorageService
}

func NewServices(memory, file storage.Storage) *Services {
	return &Services{
		Mss: NewMemStorageService(&memory),
		Fss: NewFileStorageService(&file),
	}
}
