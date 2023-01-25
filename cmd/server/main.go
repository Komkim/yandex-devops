package main

import (
	"log"
	"yandex-devops/config"
	"yandex-devops/internal/server/app"
)

func main() {
	cfg, err := config.IninServer()
	//cfg, err := config.InitFlagServer()
	if err != nil {
		log.Println(err)
	}
	app.Run(cfg)
}
