package main

import (
	"maycms/internal/interfaces"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	contentHandler := interfaces.NewContentHandler()
	server.GET("/contents", contentHandler.GetContentHandler)

	server.Run(":8000")
}
