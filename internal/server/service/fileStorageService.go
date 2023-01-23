package service

import (
	"log"
	"time"
	"yandex-devops/config"
	"yandex-devops/storage"
)

type FileStorageService struct {
	repo storage.Storage
}

func NewFileStorageService(r *storage.Storage) *FileStorageService {
	return &FileStorageService{*r}
}

func (f *FileStorageService) Start(cfg *config.Config, memStorage storage.Storage) {
	ticker := time.NewTicker(cfg.File.Interval)

	for {
		<-ticker.C
		memStorage, err := memStorage.GetAll()
		if err != nil {
			log.Println(err)
		} else {
			_, err := f.repo.SetAll(*memStorage)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (f *FileStorageService) GetAll() (*[]storage.Metrics, error) {
	return f.repo.GetAll()
}

func (f *FileStorageService) SetAll(metrics []storage.Metrics) (*[]storage.Metrics, error) {
	return f.repo.SetAll(metrics)
}
