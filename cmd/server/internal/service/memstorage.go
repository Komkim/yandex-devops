package service

import "Komkim/go-musthave-devops-tpl/cmd/server/storage"

type MemStorageService struct {
	repo storage.Storage
}

func NewMemStorageService(r storage.Storage) *MemStorageService {
	return &MemStorageService{r}
}

func (m *MemStorageService) SaveOrUpdate(metric storage.Metric) error {
	return m.repo.SetOne(metric)
}

func (m *MemStorageService) SaveOrUpdateAll(metrics []storage.Metric) error {
	return m.repo.SetAll(metrics)
}

func (m *MemStorageService) GetByKey(key string) (storage.Metric, error) {
	return m.repo.GetOne(key)
}

func (m *MemStorageService) GetAll() ([]storage.Metric, error) {
	return m.repo.GetAll()
}
