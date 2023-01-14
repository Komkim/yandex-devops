package entity

import (
	"sync"
	"yandex-devops/internal/server/storage"
)

type MemStorage struct {
	mutex     *sync.RWMutex
	storage   map[string]storage.Metric
	keySlice  []string
	typeSlice []string
}

func NewMemStorage(keySlice []string, typeSlice []string) MemStorage {
	return MemStorage{
		mutex:     &sync.RWMutex{},
		storage:   make(map[string]storage.Metric),
		keySlice:  keySlice,
		typeSlice: typeSlice,
	}
}

func (s MemStorage) GetOne(key string) (storage.Metric, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	//if !s.checkKey(key) {
	//	return storage.Metric{}, errors.New("Bad key")
	//}

	if m, ok := s.storage[key]; ok {
		return m, nil
	} else {
		return storage.Metric{}, nil
	}
}

func (s MemStorage) GetAll() ([]storage.Metric, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var metricSlice []storage.Metric
	for _, m := range s.storage {
		metricSlice = append(metricSlice, m)
	}

	return metricSlice, nil
}

func (s MemStorage) SetOne(metric storage.Metric) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	//if !s.checkKey(metric.Name) {
	//	return errors.New("Bad key")
	//}
	//
	//if !s.checkType(metric.Type) {
	//	return errors.New("Bad type")
	//}

	s.storage[metric.Name] = metric

	return nil
}

func (s MemStorage) SetAll(metric []storage.Metric) error {
	for _, m := range metric {
		if err := s.SetOne(m); err != nil {
			return err
		}
	}

	return nil
}

func (s MemStorage) checkKey(key string) bool {
	for _, v := range s.keySlice {
		if v == key {
			return true
		}
	}
	return false
}

func (s MemStorage) checkType(t string) bool {
	for _, v := range s.typeSlice {
		if v == t {
			return true
		}
	}
	return false
}
