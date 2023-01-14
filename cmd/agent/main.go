package main

import (
	"yandex-devops/config"
	"yandex-devops/internal/agent/app"
)

func main() {
	app.Run(config.Init())
}
