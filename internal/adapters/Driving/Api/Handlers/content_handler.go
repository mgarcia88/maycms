package api

import (
	DTO "maycms/internal/adapters/driving/api/DTOs/content"
	"maycms/internal/domain/entities"
	"maycms/internal/domain/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContentHandler struct {
	getAllContentsUseCase usecases.GetAllContentsUseCase
	getContentByIdUseCase usecases.GetContentByIdUseCase
	postContentUseCase    usecases.PostContentUseCase
}

func NewContentHandler(
	getAll usecases.GetAllContentsUseCase,
	getByID usecases.GetContentByIdUseCase,
	post usecases.PostContentUseCase,
) *ContentHandler {
	return &ContentHandler{
		getAllContentsUseCase: getAll,
		getContentByIdUseCase: getByID,
		postContentUseCase:    post,
	}
}

func (h *ContentHandler) HandleGetAll(c *gin.Context) {
	contents, err := h.getAllContentsUseCase.Execute()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve contents"})
		return
	}
	if len(contents) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No contents found"})
		return
	}

	var responses []DTO.ContentResponse
	for _, content := range contents {
		responses = append(responses, DTO.ContentResponse{
			ID:          content.ID,
			Title:       content.Title,
			ContentText: content.ContentText,
			Status:      content.Status,
		})
	}
	// Return the contents as JSON
	c.JSON(http.StatusOK, responses)
}

func (h *ContentHandler) HandleGetById(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	content, err := h.getContentByIdUseCase.Execute(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "No content found with the given ID")
		return
	}

	result := DTO.ContentResponse{
		ID:          content.ID,
		Title:       content.Title,
		ContentText: content.ContentText,
		Status:      content.Status,
	}

	c.JSON(http.StatusOK, result)

}

func (h *ContentHandler) HandleCreate(c *gin.Context) {
	var createContentDTO DTO.ContentRequest

	if err := c.ShouldBindJSON(&createContentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var content, err = entities.NewContent(createContentDTO.Title, createContentDTO.ContentText, createContentDTO.Status)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.postContentUseCase.Execute(content)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Content created successfully")
}
