package service

import "yandex-devops/internal/server/storage"

type Services struct {
	Storage storage.Storage
}

func NewServices(r storage.Storage) *Services {
	return &Services{NewMemStorageService(r)}
}
