package entity

import (
	"server/storage"
	"sync"
)

type MemStorage struct {
	*sync.RWMutex
	storage map[string]storage.Metric
}

func NewMemStorage() MemStorage {
	return MemStorage{
		storage: make(map[string]storage.Metric),
	}
}

func (s MemStorage) GetOne(key string) (storage.Metric, error) {
	s.RLock()
	defer s.RUnlock()

	if m, ok := s.storage[key]; ok {
		return m, nil
	} else {
		return storage.Metric{}, nil
	}
}

func (s MemStorage) GetAll() ([]storage.Metric, error) {
	s.RLock()
	defer s.RUnlock()

	var metricSlice []storage.Metric
	for _, m := range s.storage {
		metricSlice = append(metricSlice, m)
	}

	return metricSlice, nil
}

func (s MemStorage) SetOne(metric storage.Metric) error {
	s.RLock()
	defer s.RUnlock()

	s.storage[metric.Name] = metric

	return nil
}

func (s MemStorage) SetAll(metric []storage.Metric) error {
	for _, m := range metric {
		s.SetOne(m)
	}

	return nil
}
