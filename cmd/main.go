package main

import (
	"fmt"
	api "maycms/Internal/Adapters/Driving/Api/Handlers"
	"maycms/internal/adapters/driven/infra/data/postgres"
	"maycms/internal/adapters/driven/infra/data/repositories"
	"os"

	"maycms/internal/domain/application"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dsn := os.Getenv("POSTGRES_PASSWORD")
	fmt.Println(dsn)
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
