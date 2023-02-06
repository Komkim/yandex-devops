package service

import "yandex-devops/storage"

type MemStorageService struct {
	repo storage.Storage
}

func NewMemStorageService(r *storage.Storage) *MemStorageService {
	return &MemStorageService{*r}
}

func (m MemStorageService) SaveOrUpdateOne(metric storage.Metrics) (*storage.Metrics, error) {
	if metric.MType == COUNTER {
		mtr, err := m.GetByKey(metric)
		if err != nil {
			return nil, err
		}
		if mtr != nil {

			c := *mtr.Delta + *metric.Delta
			metric.Delta = &c
		}
	}

	return m.repo.SetOne(metric)
}

func (m MemStorageService) SaveOrUpdateAll(metrics []storage.Metrics) (*[]storage.Metrics, error) {
	return m.repo.SetAll(metrics)
}

func (m MemStorageService) GetByKey(metric storage.Metrics) (*storage.Metrics, error) {
	return m.repo.GetOne(metric.ID)
}

func (m MemStorageService) GetAll() (*[]storage.Metrics, error) {
	return m.repo.GetAll()
}
