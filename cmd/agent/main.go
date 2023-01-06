package main

import (
	"cmd/internal/app"
	transport "cmd/pkg"
)

func main() {

	client := transport.New()
	app.Run(client)
}
