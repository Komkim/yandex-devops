package service

import (
	"yandex-devops/storage"
)

type FileStorageService struct {
	repo storage.Storage
}

func NewFileStorageService(r *storage.Storage) *FileStorageService {
	return &FileStorageService{*r}
}

func (f *FileStorageService) GetAll() (*[]storage.Metrics, error) {
	return f.repo.GetAll()
}

func (f *FileStorageService) SetAll(metrics []storage.Metrics) (*[]storage.Metrics, error) {
	return f.repo.SetAll(metrics)
}
