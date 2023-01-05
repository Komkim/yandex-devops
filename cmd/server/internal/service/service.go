package service

import "Komkim/go-musthave-devops-tpl/cmd/server/storage"

type MemStorage interface {
	SaveOrUpdate(metric storage.Metric) error
	SaveOrUpdateAll(metrics []storage.Metric) error
	GetByKey(key string) (storage.Metric, error)
	GetAll() ([]storage.Metric, error)
}

type Services struct {
	MemStorage MemStorage
}

func NewServices(r *storage.Repositories) *Services {
	return &Services{NewMemStorageService(r.Storage)}
}
