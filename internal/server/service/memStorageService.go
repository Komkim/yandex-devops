package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"yandex-devops/storage"
)

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

func (m MemStorageService) SaveOrUpdateAll(metrics []storage.Metrics) ([]storage.Metrics, error) {
	return m.repo.SetAll(metrics)
}

func (m MemStorageService) GetByKey(metric storage.Metrics) (*storage.Metrics, error) {
	return m.repo.GetOne(metric.ID)
}

func (m MemStorageService) GetAll() ([]storage.Metrics, error) {
	return m.repo.GetAll()
}

func (m MemStorageService) CheckHash(metric storage.Metrics, key string) (bool, error) {
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

func (m MemStorageService) GenerageHash(metric storage.Metrics, key string) []byte {
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
