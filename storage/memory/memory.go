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

func (s *MemStorage) GetOne(key string) (*storage.Metrics, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if m, ok := s.storage[key]; ok {
		return &m, nil
	} else {
		return nil, nil
	}
}

func (s *MemStorage) GetAll() (*[]storage.Metrics, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var metricSlice []storage.Metrics
	for _, m := range s.storage {
		metricSlice = append(metricSlice, m)
	}

	return &metricSlice, nil
}

func (s *MemStorage) SetOne(metric storage.Metrics) (*storage.Metrics, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	s.storage[metric.ID] = metric
	r := s.storage[metric.ID]

	return &r, nil
}

func (s *MemStorage) SetAll(metric []storage.Metrics) (*[]storage.Metrics, error) {
	mm := make([]storage.Metrics, 0, len(metric))
	for _, m := range metric {
		if ss, err := s.SetOne(m); err != nil {
			return nil, err
		} else {
			mm = append(mm, *ss)
		}
	}

	return &mm, nil
}
