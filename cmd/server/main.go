package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"yandex-devops/config"
	"yandex-devops/internal/server/app"
	"yandex-devops/storage/file"
	"yandex-devops/storage/memory"
)

func main() {
	ctx, cencel := context.WithCancel(context.Background())
	defer cencel()

	//cfg, err := config.IninServer()
	cfg, err := config.InitFlagServer()
	if err != nil {
		log.Println(err)
	}

	memoryStorage := memory.NewMemStorage()
	//mm := *memoryStorage
	if memoryStorage == nil {
		log.Println("memory storage error")
		return
	}
	fileStorage, err := file.NewFileStorage(&cfg.File)
	if err != nil && fileStorage == nil {
		log.Println("file storage error")
		return
	}

	err = app.FileRestore(&cfg.File, fileStorage, memoryStorage)
	if err != nil {
		log.Println(err)
	}

	go func() {
		err = app.StartFile(ctx, &cfg.File, memoryStorage, fileStorage)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	go func() {
		err = app.StartServer(&cfg.HTTP, memoryStorage, fileStorage)
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
