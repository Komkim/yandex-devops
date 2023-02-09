package main

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
	"log"
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
	postgresql "yandex-devops/storage/postgre"
	_ "yandex-devops/storage/postgre/migrations"
)

func main() {
	ctx, cencel := context.WithCancel(context.Background())

	cfg, err := config.InitFlagServer()
	if err != nil {
		log.Println(err)
	}

	memoryStorage := memory.NewMemStorage()
	if memoryStorage == nil {
		log.Panic("memory storage error")
		return
	}

	myStorage := selectionStorage(ctx, &cfg.Server)

	s := service.NewServices(&cfg.Server, memoryStorage, myStorage)

	s.StorageService.Restore()
	go s.StorageService.Start(ctx)

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

	s.StorageService.Finish()

	defer closeStorage(&myStorage)
	defer cencel()
}

func closeStorage(storage *storage.Storage) {
	if storage != nil {
		s := *storage
		err := s.Close()
		if err != nil {
			log.Println(err)
		}
	}
}

func selectionStorage(ctx context.Context, cfg *config.Server) storage.Storage {
	switch len(cfg.DatabaseDSN) > 0 {
	case true:
		dbStorage, err := postgresql.New(ctx, cfg.DatabaseDSN)
		if err != nil {
			log.Println(err)
			return file.NewFileStorage(cfg)
		}

		db, err := sql.Open("pgx", cfg.DatabaseDSN)
		if err != nil {
			log.Println(err)
			return file.NewFileStorage(cfg)
		}

		err = goose.Up(db, "/var")
		if err != nil {
			log.Println()
			return file.NewFileStorage(cfg)
		}

		return dbStorage
	}
	return file.NewFileStorage(cfg)
}
