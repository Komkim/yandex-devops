package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"yandex-devops/storage"
)

// StorageService - сервис для работы с хранилищем
type StorageService struct {
	//repo - хранилище
	repo storage.Storage
}

// NewStorageService - создание нового сервиса
func NewStorageService(r storage.Storage) *StorageService {
	return &StorageService{r}
}

// SaveOrUpdateOne - запись новой метрики
func (m StorageService) SaveOrUpdateOne(metric storage.Metrics, key string) (*storage.Metrics, error) {
	metric, err := m.checkCounter(metric, key)
	if err != nil {
		return nil, err
	}

	return m.repo.SetOne(metric)
}

// SaveOrUpdateAll - запмсь нескольких метрик
func (m StorageService) SaveOrUpdateAll(metrics []storage.Metrics, key string) ([]storage.Metrics, error) {
	result := make([]storage.Metrics, 0, len(metrics))
	for _, mtr := range metrics {
		temp, err := m.SaveOrUpdateOne(mtr, key)
		if err != nil {
			return nil, err
		}
		result = append(result, *temp)
	}
	return result, nil
}

// GetByKey - получение одноий метрики
func (m StorageService) GetByKey(metric storage.Metrics) (*storage.Metrics, error) {
	return m.repo.GetOne(metric.ID)
}

// GetAll - получение всех метрик
func (m StorageService) GetAll() ([]storage.Metrics, error) {
	return m.repo.GetAll()
}

// CheckHash - проверка хэша
func (m StorageService) CheckHash(metric storage.Metrics, key string) (bool, error) {
	if len(key) <= 0 {
		return true, nil
	}

	if len(metric.Hash) <= 0 {
		return true, nil
	}

	h1, err := hex.DecodeString(metric.Hash)
	if err != nil {
		return false, err
	}

	h2 := m.GenerageHash(metric, key)

	return hmac.Equal(h1, h2), nil
}

// GenerageHash - генерация хэша
func (m StorageService) GenerageHash(metric storage.Metrics, key string) []byte {
	var data []byte
	switch metric.MType {
	case COUNTER:
		if metric.Delta != nil && len(metric.ID) > 0 && len(metric.MType) > 0 {
			data = []byte(fmt.Sprintf("%s:%s:%d", metric.ID, metric.MType, *metric.Delta))
		}
	case GAUGE:
		if metric.Value != nil && len(metric.ID) > 0 && len(metric.MType) > 0 {
			data = []byte(fmt.Sprintf("%s:%s:%f", metric.ID, metric.MType, *metric.Value))
		}
	}

	h := hmac.New(sha256.New, []byte(key))
	h.Write(data)
	return h.Sum(nil)
}

// checkCounter - проверка счетчика
func (m StorageService) checkCounter(metric storage.Metrics, key string) (storage.Metrics, error) {
	if metric.MType == COUNTER {
		mtr, err := m.repo.GetOne(metric.ID)
		if err != nil {
			return metric, err
		}
		if mtr != nil && mtr.Delta != nil && metric.Delta != nil {
			c := *mtr.Delta + *metric.Delta
			metric.Delta = &c
			metric.Hash = hex.EncodeToString(m.GenerageHash(metric, key))
		}
	}
	if len(metric.Hash) <= 0 && len(key) > 0 {
		metric.Hash = hex.EncodeToString(m.GenerageHash(metric, key))
	}
	return metric, nil
}
