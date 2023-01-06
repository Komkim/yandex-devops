package app

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"server/internal/entity"
	router "server/internal/http"
	"server/internal/service"
	server "server/pkg"
	"server/storage"
	"syscall"
)

func Run() {

	rep := storage.NewRepositories(entity.NewMemStorage(keyInit(), typeInit()))
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

func keyInit() []string {
	return []string{
		"Alloc",
		"BuckHashSys",
		"Frees",
		"GCCPUFraction",
		"GCSys",
		"HeapAlloc",
		"HeapIdle",
		"HeapInuse",
		"HeapObjects",
		"HeapReleased",
		"HeapSys",
		"LastGC",
		"Lookups",
		"MCacheInuse",
		"MCacheSys",
		"MSpanInuse",
		"MSpanSys",
		"Mallocs",
		"NextGC",
		"NumForcedGC",
		"NumGC",
		"OtherSys",
		"PauseTotalNs",
		"StackInuse",
		"StackSys",
		"Sys",
		"TotalAlloc",

		"PollCount",
		"RandomValue",
	}
}

func typeInit() []string {
	return []string{
		"gauge",
		"counter",
	}
}
