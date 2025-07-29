package main

import (
	"fmt"
	"maycms/internal/adapters/driven/infra/data/postgres"
	"maycms/internal/adapters/driven/infra/data/repositories"
	api "maycms/internal/adapters/driving/api/handlers"
	"maycms/internal/domain/usecases"
	"os"

	"maycms/internal/application"

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
	categoryRepo := repositories.NewCategoryRepository(db)

	getAllContentsUseCase := usecases.NewGetAllContentsUseCase(*contentRepo)
	getContentByIdUseCase := usecases.NewGetContentByIdUseCase(*contentRepo)
	postContentUseCase := usecases.NewPostContentUseCase(*contentRepo)

	contentService := application.NewContentService(
		*getAllContentsUseCase,
		*getContentByIdUseCase,
		*postContentUseCase)

	categoryService := application.NewCategoryService(*categoryRepo)

	contentHandler := api.NewContentHandler(*contentService)
	categoryHandler := api.NewCategoryHandler(*categoryService)

	server.GET("/contents", contentHandler.GetContentHandler)
	server.GET("/contents/:id", contentHandler.GetContentByIDHandler)
	server.POST("/contents", contentHandler.CreateContentHandler)

	server.POST("/categories", categoryHandler.CreateCategoryHandler)

	server.Run(":8080")
}
