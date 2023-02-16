package service

import (
	"context"
	"log"
	"time"
	"yandex-devops/config"
	"yandex-devops/storage"
)

type StorageService struct {
	repo          storage.Storage
	cfg           *config.Server
	memoryService *MemStorageService
}

func NewStorageService(cfg *config.Server, r *storage.Storage, memoryService *MemStorageService) *StorageService {
	return &StorageService{cfg: cfg, repo: *r, memoryService: memoryService}
}

func (s *StorageService) GetAll() ([]storage.Metrics, error) {
	return s.repo.GetAll()
}

func (s *StorageService) SetAll(metrics []storage.Metrics) ([]storage.Metrics, error) {
	return s.repo.SetAll(metrics)
}

func (s *StorageService) Restore() {
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

	_, err = s.memoryService.SaveOrUpdateAll(metrics)
	if err != nil {
		return
	}
}

func (s *StorageService) Start(ctx context.Context) {
	ticker := time.NewTicker(s.cfg.FileInterval)

	for {
		select {
		case <-ticker.C:
			if err := s.record(); err != nil {
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}

func (s *StorageService) Finish() {
	if err := s.record(); err != nil {
		log.Println(err)
	}
}

func (s *StorageService) record() error {
	metrics, err := s.memoryService.GetAll()

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
