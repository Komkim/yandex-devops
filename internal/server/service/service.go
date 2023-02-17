package service

import (
	"yandex-devops/storage"
)

const GAUGE = "gauge"
const COUNTER = "counter"

type Services struct {
	StorageService *StorageService
}

func NewServices(storage storage.Storage) *Services {
	return &Services{
		StorageService: NewStorageService(storage),
	}
}
