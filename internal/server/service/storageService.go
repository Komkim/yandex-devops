package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"yandex-devops/storage"
)

type StorageService struct {
	repo storage.Storage
}

func NewStorageService(r storage.Storage) *StorageService {
	return &StorageService{r}
}

func (m StorageService) SaveOrUpdateOne(metric storage.Metrics, key string) (*storage.Metrics, error) {
	if metric.MType == COUNTER {
		mtr, err := m.GetByKey(metric)
		if err != nil {
			return nil, err
		}
		if mtr != nil {

			c := *mtr.Delta + *metric.Delta
			metric.Delta = &c
			metric.Hash = hex.EncodeToString(m.GenerageHash(metric, key))
		}
	}

	return m.repo.SetOne(metric)
}

func (m StorageService) SaveOrUpdateAll(metrics []storage.Metrics) ([]storage.Metrics, error) {
	return m.repo.SetAll(metrics)
}

func (m StorageService) GetByKey(metric storage.Metrics) (*storage.Metrics, error) {
	return m.repo.GetOne(metric.ID)
}

func (m StorageService) GetAll() ([]storage.Metrics, error) {
	return m.repo.GetAll()
}

func (m StorageService) CheckHash(metric storage.Metrics, key string) (bool, error) {
	if len(key) <= 0 {
		return true, nil
	}

	h1, err := hex.DecodeString(metric.Hash)
	if err != nil {
		return false, err
	}

	h2 := m.GenerageHash(metric, key)

	return hmac.Equal(h1, h2), nil
}

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
