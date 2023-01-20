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
	"yandex-devops/storage/memory"
)

func Run(config *config.Config) {

	srv := service.NewServices(memory.NewMemStorage())
	r := router.NewRouter(srv)
	s := server.NewServer(config, r.Init())

	go func() {
		if err := s.Start(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
