package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"yandex-devops/config"
	"yandex-devops/internal/server/app"
	router "yandex-devops/internal/server/http"
	"yandex-devops/internal/server/server"
	"yandex-devops/internal/server/service"
	"yandex-devops/storage/file"
	"yandex-devops/storage/memory"
)

func main() {
	ctx, cencel := context.WithCancel(context.Background())
	defer cencel()

	cfg, err := config.InitFlagServer()
	if err != nil {
		log.Println(err)
	}

	memoryStorage := memory.NewMemStorage()
	if memoryStorage == nil {
		log.Println("memory storage error")
		return
	}
	fileStorage, err := file.NewFileStorage(&cfg.Server)
	if err != nil && fileStorage == nil {
		log.Println("file storage error")
		return
	}

	myFile := app.NewMyFile(ctx, &cfg.Server, memoryStorage, fileStorage)

	go myFile.Restore()
	go myFile.Start()
	defer myFile.Finish()

	go func() {
		err = startServer(&cfg.HTTP, memoryStorage, fileStorage)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	defer func() {
		if fileStorage != nil {
			err := fileStorage.Close()
			if err != nil {
				log.Println(err)
			}
		}
	}()
}

func startServer(cfg *config.HTTP, memStorage *memory.MemStorage, fileStorage *file.FileStorage) error {
	srv := service.NewServices(memStorage, fileStorage)
	r := router.NewRouter(srv)
	s := server.NewServer(cfg, r.Init())

	return s.Start()
}
