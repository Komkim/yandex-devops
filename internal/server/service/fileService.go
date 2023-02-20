package service

import (
	"context"
	"log"
	"time"
	"yandex-devops/config"
	"yandex-devops/storage"
)

type FileService struct {
	repo           storage.Storage
	cfg            *config.Server
	storageService *StorageService
}

func NewFileService(cfg *config.Server, r storage.Storage, storageService *StorageService) *FileService {
	return &FileService{cfg: cfg, repo: r, storageService: storageService}
}

func (s *FileService) GetAll() ([]storage.Metrics, error) {
	return s.repo.GetAll()
}

func (s *FileService) SetAll(metrics []storage.Metrics) ([]storage.Metrics, error) {
	return s.repo.SetAll(metrics)
}

func (s *FileService) restore() {
	if !s.cfg.FileRestore {
		return
	}
	if s == nil {
		return
	}

	metrics, err := s.GetAll()
	if err != nil {
		return
	}

	_, err = s.storageService.SaveOrUpdateAll(metrics, s.cfg.Key)
	if err != nil {
		return
	}
}

func (s *FileService) Start(ctx context.Context) {
	s.restore()

	ticker := time.NewTicker(s.cfg.FileInterval)

n:
	for {
		select {
		case <-ticker.C:
			if err := s.record(); err != nil {
				continue
			}
		case <-ctx.Done():
			break n
		}
	}

	defer s.finish()
}

func (s *FileService) finish() {
	if err := s.record(); err != nil {
		log.Println(err)
	}
}

func (s *FileService) record() error {
	metrics, err := s.storageService.GetAll()

	if err != nil {
		return err
	} else {
		_, err := s.SetAll(metrics)
		if err != nil {
			return err
		}
	}
	return nil
}
