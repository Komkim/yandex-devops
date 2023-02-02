package app

import (
	"context"
	"errors"
	"fmt"
	"time"
	"yandex-devops/config"
	router "yandex-devops/internal/server/http"
	"yandex-devops/internal/server/server"
	"yandex-devops/internal/server/service"
	"yandex-devops/storage/file"
	"yandex-devops/storage/memory"
)

func StartServer(cfg *config.HTTP, memStorage *memory.MemStorage, fileStorage *file.FileStorage) error {
	srv := service.NewServices(memStorage, fileStorage)
	r := router.NewRouter(srv)
	s := server.NewServer(cfg, r.Init())

	return s.Start()
}

func FileRestore(cfg *config.File, fileStorage *file.FileStorage, memStorage *memory.MemStorage) error {
	if !cfg.Restore {
		return nil
	}
	if fileStorage == nil {
		return errors.New("file storage nil")
	}

	metrics, err := fileStorage.GetAll()
	if err != nil {
		return err
	}

	_, err = memStorage.SetAll(*metrics)
	if err != nil {
		return err
	}

	return nil
}

func StartFile(ctx context.Context, config *config.File, memStorage *memory.MemStorage, fileStorage *file.FileStorage) error {
	ticker := time.NewTicker(config.Interval)
	ms := *memStorage
	fs := *fileStorage

	for {
		select {
		case <-ticker.C:
			metrics, err := ms.GetAll()
			if err != nil {
				return err
			} else {
				_, err := fs.SetAll(*metrics)
				if err != nil {
					return err
				}
			}
		case <-ctx.Done():
			return fmt.Errorf("context canceled")
		}
	}
}

func FinishFile() {

}
