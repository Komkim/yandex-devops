package main

import (
	"github.com/caarlos0/env/v6"
	"log"
	"yandex-devops/config"
	"yandex-devops/internal/agent/app"
)

func main() {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		log.Println(err)
	}
	app.Run(&cfg)

}
