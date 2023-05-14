// Внетренние сервисы сервера
package service

import (
	"context"
	"log"
	"time"
	"yandex-devops/config"
	"yandex-devops/storage"
)

// FileService - сервис для работы с файлами
type FileService struct {
	//repo - внутренний репозиторий сервиса
	repo storage.Storage
	//cfg - конфиг
	cfg *config.Server
	//storageService - репозиторий для работы с базой
	storageService *StorageService
}

// NewFileService - создание нового сервиса для работы с файлами
func NewFileService(cfg *config.Server, r storage.Storage, storageService *StorageService) *FileService {
	return &FileService{cfg: cfg, repo: r, storageService: storageService}
}

// GetAll - получение всех метрик из файла
func (s *FileService) GetAll() ([]storage.Metrics, error) {
	return s.repo.GetAll()
}

// SetAll - запись нескольких метрик в файл
func (s *FileService) SetAll(metrics []storage.Metrics) ([]storage.Metrics, error) {
	return s.repo.SetAll(metrics)
}

// restore - восстановление записей из файла
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

// Start - запуск сервиса для работы с файлами
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

// finish - завершение работы
func (s *FileService) finish() {
	if err := s.record(); err != nil {
		log.Println(err)
	}
}

// record - запись метрик в файл
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
