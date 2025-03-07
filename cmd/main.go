package main

import (
	"maycms/internal/adapters/driven/infra/data/postgres"
	"maycms/internal/adapters/driven/infra/data/repositories"
	api "maycms/internal/adapters/drivers/api/handlers"
	"maycms/internal/domain/application"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db := postgres.NewPostgresDB()
	contentRepo := repositories.NewContentRepository(db)
	contentService := application.NewContentService(*contentRepo)

	contentHandler := api.NewContentHandler(*contentService)
	server.GET("/contents", contentHandler.GetContentHandler)
	server.GET("/contents/:id", contentHandler.GetContentByIDHandler)
	server.POST("/contents", contentHandler.CreateContentHandler)

	server.Run(":8000")
}
