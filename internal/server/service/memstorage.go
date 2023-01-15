package service

import "yandex-devops/internal/server/storage"

type MemStorageService struct {
	repo storage.Storage
}

func NewMemStorageService(r storage.Storage) MemStorageService {
	return MemStorageService{r}
}

func (m MemStorageService) SetOne(metric storage.Metric) error {
	return m.repo.SetOne(metric)
}

func (m MemStorageService) SetAll(metrics []storage.Metric) error {
	return m.repo.SetAll(metrics)
}

func (m MemStorageService) GetOne(key string) (storage.Metric, error) {
	return m.repo.GetOne(key)
}

func (m MemStorageService) GetAll() ([]storage.Metric, error) {
	return m.repo.GetAll()
}
