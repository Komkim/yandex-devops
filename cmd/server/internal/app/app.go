package app

import (
	"Komkim/go-musthave-devops-tpl/cmd/server/internal/entity"
	router "Komkim/go-musthave-devops-tpl/cmd/server/internal/http"
	"Komkim/go-musthave-devops-tpl/cmd/server/internal/service"
	server "Komkim/go-musthave-devops-tpl/cmd/server/pkg"
	"Komkim/go-musthave-devops-tpl/cmd/server/storage"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() {

	rep := storage.NewRepositories(entity.NewMemStorage())
	srv := service.NewServices(rep)

	r := router.NewRouter(srv)
	s := server.NewServer(r.Init())

	go func() {
		if err := s.Start(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	s.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
