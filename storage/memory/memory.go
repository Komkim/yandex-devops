package memory

import (
	"sync"
	"yandex-devops/storage"
)

type MemStorage struct {
	mutex   *sync.RWMutex
	storage map[string]storage.Metrics
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		mutex:   &sync.RWMutex{},
		storage: make(map[string]storage.Metrics),
	}
}

func (s *MemStorage) GetOne(key string) (storage.Metrics, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if m, ok := s.storage[key]; ok {
		return m, nil
	} else {
		return storage.Metrics{}, nil
	}
}

func (s *MemStorage) GetAll() ([]storage.Metrics, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var metricSlice []storage.Metrics
	for _, m := range s.storage {
		metricSlice = append(metricSlice, m)
	}

	return metricSlice, nil
}

func (s *MemStorage) SetOne(metric storage.Metrics) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	s.storage[metric.ID] = metric

	return nil
}

func (s *MemStorage) SetAll(metric []storage.Metrics) error {
	for _, m := range metric {
		if err := s.SetOne(m); err != nil {
			return err
		}
	}

	return nil
}
