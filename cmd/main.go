package main

import (
	"maycms/internal/adapters/drivers/api"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	contentHandler := api.NewContentHandler()
	server.GET("/contents", contentHandler.GetContentHandler)
	server.GET("/contents/:id", contentHandler.GetContentByIDHandler)

	server.Run(":8000")
}
