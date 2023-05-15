// server - модуль для обработки и хранения метрик
package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"yandex-devops/config"
	router "yandex-devops/internal/server/http"
	"yandex-devops/internal/server/server"
	"yandex-devops/internal/server/service"
	"yandex-devops/storage"
	"yandex-devops/storage/file"
	"yandex-devops/storage/memory"
	postgresql "yandex-devops/storage/postgre"
	_ "yandex-devops/storage/postgre/migrations"

	"github.com/pressly/goose/v3"
)

func main() {
	ctx, cencel := context.WithCancel(context.Background())

	cfg, err := config.InitFlagServer()
	if err != nil {
		log.Println(err)
	}

	log.Println(cfg)

	myStorage := selectionStorage(ctx, &cfg.Server)

	s := service.NewServices(myStorage)

	if len(cfg.Server.DatabaseDSN) <= 0 {

		fileStorage := file.NewFileStorage(&cfg.Server)
		if fileStorage == nil {
			log.Println("file storage error")

		}
		fileService := service.NewFileService(&cfg.Server, fileStorage, s.StorageService)
		go fileService.Start(ctx)
	}

	r := router.NewRouter(&cfg.Server, s)
	srv := server.NewServer(&cfg.HTTP, r.Init())

	go srv.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.GetServer().Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	defer closeStorage(myStorage)
	defer cencel()
}

func closeStorage(storage storage.Storage) {
	if storage != nil {
		s := storage
		err := s.Close()
		if err != nil {
			log.Println(err)
		}
	}
}

func selectionStorage(ctx context.Context, cfg *config.Server) storage.Storage {
	if len(cfg.DatabaseDSN) <= 0 {
		return memory.NewMemStorage()
	}

	dbStorage, err := postgresql.New(ctx, cfg.DatabaseDSN)
	if err != nil {
		log.Println(err)
		return memory.NewMemStorage()
	}

	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		log.Println(err)
		return memory.NewMemStorage()
	}

	err = goose.Up(db, "/var")
	if err != nil {
		log.Println()
		return memory.NewMemStorage()
	}

	return dbStorage
}
