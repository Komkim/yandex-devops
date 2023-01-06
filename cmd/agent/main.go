package main

import (
	"agent/internal/app"
	transport "agent/pkg"
)

func main() {

	client := transport.New()
	app.Run(client)
}
