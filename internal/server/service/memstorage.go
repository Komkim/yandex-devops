package service

import "yandex-devops/storage"

type MemStorageService struct {
	repo storage.Storage
}

func NewMemStorageService(r *storage.Storage) *MemStorageService {
	return &MemStorageService{*r}
}

func (m MemStorageService) SaveOrUpdateOne(metric storage.Metrics) error {
	if metric.MType == COUNTER {
		mtr, err := m.GetByKey(metric)
		if err != nil {
			return err
		}
		if mtr != (storage.Metrics{}) {
			c1 := mtr.Delta
			c2 := metric.Delta
			c := *c1 + *c2
			metric.Delta = &c
		}
	}

	return m.repo.SetOne(metric)
}

func (m MemStorageService) SaveOrUpdateAll(metrics []storage.Metrics) error {
	return m.repo.SetAll(metrics)
}

func (m MemStorageService) GetByKey(metric storage.Metrics) (storage.Metrics, error) {
	mm, err := m.repo.GetOne(metric.ID)
	if err != nil {
		return storage.Metrics{}, err
	}
	if mm != (storage.Metrics{}) && mm.MType != metric.MType {
		return mm, nil
	}
	return mm, nil
}

func (m MemStorageService) GetAll() ([]storage.Metrics, error) {
	return m.repo.GetAll()
}
