package main

import (
	"log"
	"yandex-devops/config"
	"yandex-devops/internal/agent/app"
)

func main() {

	cfg, err := config.IninServer()
	if err != nil {
		log.Println(err)
	}
	app.Run(cfg)

}
