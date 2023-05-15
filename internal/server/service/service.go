package service

import (
	"yandex-devops/storage"
)

// типы метрик
const (
	//GAUGE - число с плавающей точкой
	GAUGE = "gauge"
	//COUNTER - счетчик
	COUNTER = "counter"
)

// Services - внутренний сервис
type Services struct {
	//StorageService - сервис для работы с хранилищем
	StorageService *StorageService
}

// NewServices - создание нового сервиса
func NewServices(storage storage.Storage) *Services {
	return &Services{
		StorageService: NewStorageService(storage),
	}
}
