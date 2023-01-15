package app

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"yandex-devops/config"
	"yandex-devops/internal/server/entity"
	router "yandex-devops/internal/server/http"
	"yandex-devops/internal/server/server"
	"yandex-devops/internal/server/service"
)

func Run(config *config.Config) {

	srv := service.NewServices(entity.NewMemStorage(keyInit(), typeInit()))

	r := router.NewRouter(srv)
	s := server.NewServer(config, r.Init())

	go func() {
		if err := s.Start(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error occurred while running http server: %s\n", err.Error())
		}
	}()

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
