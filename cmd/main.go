package main

import (
	"fmt"
	"maycms/internal/adapters/driven/infra/data/postgres"
	"maycms/internal/adapters/driven/infra/data/repositories"
	api "maycms/internal/adapters/driving/api/handlers"
	"maycms/internal/domain/usecases"
	"os"

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

	getAllContentsUseCase := usecases.NewGetAllContentsUseCase(*contentRepo)
	getContentByIdUseCase := usecases.NewGetContentByIdUseCase(*contentRepo)
	postContentUseCase := usecases.NewPostContentUseCase(*contentRepo)

	contentHandler := api.NewContentHandler(
		*getAllContentsUseCase,
		*getContentByIdUseCase,
		*postContentUseCase)

	server.GET("/contents", contentHandler.HandleGetAll)
	server.GET("/contents/:id", contentHandler.HandleGetById)
	server.POST("/contents", contentHandler.HandleCreate)

	server.Run(":8080")
}
