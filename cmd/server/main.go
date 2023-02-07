package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
		log.Panic("memory storage error")
		return
	}
	fileStorage, err := file.NewFileStorage(&cfg.Server)
	if err != nil && fileStorage == nil {
		log.Fatal("file storage error")
		return
	}

	s := service.NewServices(memoryStorage, fileStorage)

	myFile := app.NewMyFile(ctx, &cfg.Server, s)

	go myFile.Restore()
	go myFile.Start()

	r := router.NewRouter(&cfg.Server, s)
	srv := server.NewServer(&cfg.HTTP, r.Init())

	go srv.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.GetServer().Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	myFile.Finish()

	defer func() {
		if fileStorage != nil {
			err := fileStorage.Close()
			if err != nil {
				log.Println(err)
			}
		}
	}()
}
