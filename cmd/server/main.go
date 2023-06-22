// server - модуль для обработки и хранения метрик
package main

import (
	"context"
	"database/sql"
	"fmt"
	//"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
	"yandex-devops/config"
	mygrpc "yandex-devops/internal/server/grpc"
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

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	fmt.Printf("Build version: %s", buildVersion)
	fmt.Println()
	fmt.Printf("Build date: %s", buildDate)
	fmt.Println()
	fmt.Printf("Build commit: %s", buildCommit)
	fmt.Println()

	ctx, cancel := context.WithCancel(context.Background())
	//g, gCtx := errgroup.WithContext(ctx)

	quit := make(chan os.Signal, 1)
	//g.Go(func() error {
	go func() {
		//quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

		//<-quit
		//cancel()
		//return nil
	}()

	cfg, err := config.InitFlagServer()
	if err != nil {
		log.Println(err)
	}

	myStorage := selectionStorage(ctx, cfg)

	s := service.NewServices(myStorage)

	if len(cfg.DatabaseDSN) <= 0 {

		fileStorage := file.NewFileStorage(cfg)
		if fileStorage == nil {
			log.Println("file storage error")

		}
		fileService := service.NewFileService(cfg, fileStorage, s.StorageService)
		go fileService.Start(ctx)
	}

	r := router.NewRouter(cfg, s)
	srv := server.NewServer(cfg, r.Init())

	grpcR := mygrpc.NewRouter(cfg, s)
	grpcSrv := server.NewGrpcServer(cfg, grpcR)

	//g.Go(func() error {
	//	return srv.Start()
	//})

	go func() {
		srv.Start()
	}()

	//g.Go(func() error {
	//	return grpcSrv.Start()
	//})

	go func() {
		grpcSrv.Start()
	}()

	//g.Go(func() error {
	//	<-gCtx.Done()
	//	return srv.GetServer().Shutdown(context.Background())
	//})

	go func() {
		<-ctx.Done()
		srv.GetServer().Shutdown(context.Background())
	}()

	//if err := g.Wait(); err != nil {
	//	fmt.Printf("exit reason: %s \n", err)
	//}

	<-quit
	cancel()

	defer closeStorage(myStorage)
	//defer cancel()
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
