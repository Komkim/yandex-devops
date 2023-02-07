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
	fileStorage := file.NewFileStorage(&cfg.Server)
	if fileStorage == nil {
		log.Fatal("file storage error")
		return
	}

	s := service.NewServices(memoryStorage, fileStorage)

	myFile := app.NewMyFile(&cfg.Server, s)

	myFile.Restore()
	go myFile.Start(ctx)

	r := router.NewRouter(&cfg.Server, s)
	srv := server.NewServer(&cfg.HTTP, r.Init())

	go srv.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx2, cancel2 := context.WithTimeout(ctx, 5*time.Second)
	defer cancel2()
	if err := srv.GetServer().Shutdown(ctx2); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	myFile.Finish()

	defer closeFileStorage(fileStorage)
}

func closeFileStorage(fileStorage *file.FileStorage) {
	if fileStorage != nil {
		err := fileStorage.Close()
		if err != nil {
			log.Println(err)
		}
	}
}
