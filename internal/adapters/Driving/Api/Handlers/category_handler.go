package api

import (
	DTO "maycms/internal/adapters/driving/api/DTOs"
	application "maycms/internal/application"
	"maycms/internal/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service application.CategoryService
}

func NewCategoryHandler(s application.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h CategoryHandler) CreateCategoryHandler(c *gin.Context) {
	var createContentDTO DTO.CategoryRequestBody

	if err := c.ShouldBindJSON(&createContentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var category, err = entities.NewCategory(createContentDTO.Title, createContentDTO.Description)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateCategory(*category)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Categoria inserida com sucesso")

}
