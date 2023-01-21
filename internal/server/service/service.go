package service

import "yandex-devops/storage"

const GAUGE = "gauge"
const COUNTER = "counter"

type Services struct {
	Mss *MemStorageService
}

func NewServices(r storage.Storage) *Services {
	return &Services{NewMemStorageService(&r)}
}
