package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"yandex-devops/config"
	router "yandex-devops/internal/server/http"
	"yandex-devops/internal/server/server"
	"yandex-devops/internal/server/service"
	"yandex-devops/storage"
	"yandex-devops/storage/file"
	"yandex-devops/storage/memory"
)

func Run(config *config.Config) {

	memoryStorage := memory.NewMemStorage()
	fileStorage, err := file.NewFileStorage(&config.File)
	if err != nil {
		log.Println("file storage error")
	}
	srv := service.NewServices(memoryStorage, fileStorage)
	r := router.NewRouter(srv)
	s := server.NewServer(&config.HTTP, r.Init())

	go func() {
		if err := s.Start(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	if fileStorage != nil {
		go func(memStorage storage.Storage) {
			if config.File.Restore {
				metrics, err := srv.Fss.GetAll()
				if err != nil {
					log.Println(metrics)
				}
				if metrics != nil {
					_, err = srv.Fss.SetAll(*metrics)
					if err != nil {
						log.Println(err)
					}
				}
			}
			srv.Fss.Start(config, memoryStorage)
		}(memoryStorage)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if fileStorage != nil {
		metrics, err := memoryStorage.GetAll()
		if err != nil {
			log.Println(err)

		}
		_, err = fileStorage.SetAll(*metrics)
		if err != nil {
			log.Println(err)

		}
		err = fileStorage.Close()
		if err != nil {
			log.Println(err)
		}
	}
}

func Start(config *config.Config) {
	ctx, cencel := context.WithCancel(context.Background())
	memoryStorage := memory.NewMemStorage()
	if memoryStorage == nil {
		log.Println("memory storage error")
		return
	}
	fileStorage, err := file.NewFileStorage(&config.File)
	if err != nil && fileStorage == nil {
		log.Println("file storage error")
		return
	}

	//err = FileRestore(&config.File, memoryStorage, fileStorage)
	//if err != nil {
	//	log.Println(err)
	//}

	go func() {
		err = StartFile(ctx, &config.File, memoryStorage, fileStorage)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	err = StartServer(&config.HTTP, memoryStorage, fileStorage)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		if fileStorage != nil {
			err := fileStorage.Close()
			if err != nil {
				log.Println(err)
			}
		}
	}()
	defer cencel()
}

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
