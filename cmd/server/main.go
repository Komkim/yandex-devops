package main

import (
	"yandex-devops/config"
	"yandex-devops/internal/server/app"
)

func main() {
	app.Run(config.Init())
}
