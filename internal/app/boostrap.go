package app

import (
	"maycms/internal/adapters/driven/infra/data/postgres"
	"maycms/internal/adapters/driven/infra/data/repositories"
	api "maycms/internal/adapters/driving/api/handlers"
	"maycms/internal/domain/usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupServer() *gin.Engine {
	// Load environment variables
	godotenv.Load()

	// Initialize DB (Could pass dsn here if needed)
	db := postgres.NewPostgresDB()
	contentRepo := repositories.NewContentRepository(db)

	// Initialize use cases (domain logic)
	getAllContents := usecases.NewGetAllContentsUseCase(*contentRepo)
	getContentById := usecases.NewGetContentByIdUseCase(*contentRepo)
	postContent := usecases.NewPostContentUseCase(*contentRepo)

	// Initialize API handlers (driving adapters)
	contentHandler := api.NewContentHandler(
		*getAllContents,
		*getContentById,
		*postContent,
	)

	// Setup Gin server and routes
	server := gin.Default()
	server.GET("/contents", contentHandler.HandleGetAll)
	server.GET("/contents/:id", contentHandler.HandleGetById)
	server.POST("/contents", contentHandler.HandleCreate)

	return server
}
