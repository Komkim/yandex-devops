// Работа с хранилищем в памяти
package memory

import (
	"sync"
	"yandex-devops/storage"
)

// MemStorage - хранилище в памяти приложения
type MemStorage struct {
	//mutex - мютекс
	mutex *sync.RWMutex
	//storage - хранилище
	storage map[string]storage.Metrics
}

// NewMemStorage - создние нового хранилища
func NewMemStorage() *MemStorage {
	return &MemStorage{
		mutex:   &sync.RWMutex{},
		storage: make(map[string]storage.Metrics),
	}
}

// GetOne - получение метрики
func (s *MemStorage) GetOne(key string) (*storage.Metrics, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if m, ok := s.storage[key]; ok {
		return &m, nil
	} else {
		return nil, nil
	}
}

// GetAll - получение всех метрик
func (s *MemStorage) GetAll() ([]storage.Metrics, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var metricSlice []storage.Metrics
	for _, m := range s.storage {
		metricSlice = append(metricSlice, m)
	}

	return metricSlice, nil
}

// SetOne - запись одной метрики
func (s *MemStorage) SetOne(metric storage.Metrics) (*storage.Metrics, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	//m := metric

	//s.storage[metric.ID] = m
	s.storage[metric.ID] = metric
	r := s.storage[metric.ID]

	return &r, nil
}

// SetAll - запись нескольих метрик
func (s *MemStorage) SetAll(metric []storage.Metrics) ([]storage.Metrics, error) {
	mm := make([]storage.Metrics, 0, len(metric))
	for _, m := range metric {
		if ss, err := s.SetOne(m); err != nil {
			return nil, err
		} else {
			mm = append(mm, *ss)
		}
	}

	return mm, nil
}

// Close - завершение работы с хранилищем
func (s *MemStorage) Close() error {
	return nil
}

// Cleaning - очистка памяти
func (s *MemStorage) Cleaning() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for k := range s.storage {
		delete(s.storage, k)
	}
}
