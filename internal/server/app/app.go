package app

import (
	"errors"
	"log"
	"net/http"
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
)

func Run(config *config.Config) {

	memoryStorage := memory.NewMemStorage()
	fileStorage, err := file.NewFileStorage(config.File.Path)
	defer fileStorage.Close()
	if err != nil {
		log.Println("file storage error")
	}
	srv := service.NewServices(memoryStorage, fileStorage)
	r := router.NewRouter(srv)
	s := server.NewServer(config, r.Init())

	go func() {
		if err := s.Start(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("error occurred while running http server: %s\n", err.Error())
		}
	}()

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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

}
