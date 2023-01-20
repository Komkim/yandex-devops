package service

import "yandex-devops/storage"

type Services struct {
	Mss *MemStorageService
}

func NewServices(r storage.Storage) *Services {
	return &Services{NewMemStorageService(&r)}
}
