package main

import (
	"maycms/internal/app"
)

func main() {
	// Setup and start the HTTP server
	server := app.SetupServer()
	if err := server.Run(":8080"); err != nil {
		panic(err) // ou log.Fatal(err)
	}
}
