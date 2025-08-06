package main

import (
	"maycms/internal/app"
)

func main() {
	// Setup and start the HTTP server
	server := app.SetupServer()
	server.Run(":8080")
}
