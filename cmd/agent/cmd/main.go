package main

import (
	"Komkim/go-musthave-devops-tpl/cmd/agent/internal/app"
	transport "Komkim/go-musthave-devops-tpl/cmd/agent/pkg"
)

func main() {

	client := transport.New()
	app.Run(client)
}
