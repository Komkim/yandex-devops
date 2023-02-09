package service

//
//import (
//	"context"
//	"log"
//	"time"
//	"yandex-devops/config"
//	"yandex-devops/storage"
//)
//
//type FileStorageService struct {
//	repo          storage.Storage
//	cfg           *config.Server
//	memoryService *MemStorageService
//}
//
//func NewFileStorageService(cfg *config.Server, r *storage.Storage, memoryService *MemStorageService) *FileStorageService {
//	return &FileStorageService{cfg: cfg, repo: *r, memoryService: memoryService}
//}
//
//func (f *FileStorageService) GetAll() (*[]storage.Metrics, error) {
//	return f.repo.GetAll()
//}
//
//func (f *FileStorageService) SetAll(metrics []storage.Metrics) (*[]storage.Metrics, error) {
//	return f.repo.SetAll(metrics)
//}
//
//func (f *FileStorageService) Restore() {
//	if !f.cfg.FileRestore {
//		return
//	}
//	if f == nil {
//		return
//	}
//
//	metrics, err := f.GetAll()
//	if err != nil {
//		return
//	}
//
//	_, err = f.memoryService.SaveOrUpdateAll(*metrics)
//	if err != nil {
//		return
//	}
//}
//
//func (f *FileStorageService) Start(ctx context.Context) {
//	ticker := time.NewTicker(f.cfg.FileInterval)
//
//	for {
//		select {
//		case <-ticker.C:
//			if err := f.recordFile(); err != nil {
//				continue
//			}
//		case <-ctx.Done():
//			return
//		}
//	}
//}
//
//func (f *FileStorageService) Finish() {
//	if err := f.recordFile(); err != nil {
//		log.Println(err)
//	}
//}
//
//func (f *FileStorageService) recordFile() error {
//	metrics, err := f.memoryService.GetAll()
//
//	if err != nil {
//		return err
//	} else {
//		_, err := f.SetAll(*metrics)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
