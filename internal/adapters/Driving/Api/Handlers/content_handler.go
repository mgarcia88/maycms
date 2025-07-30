package api

import (
	DTO "maycms/internal/adapters/driving/api/DTOs"
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
	// Return the contents as JSON
	c.JSON(http.StatusOK, contents)
}

func (h *ContentHandler) HandleGetById(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	result, _ := h.getContentByIdUseCase.Execute(id)

	if result.ContentText == "" {
		c.JSON(http.StatusNotFound, "Conteúdo não existe")
		return
	}
	c.JSON(http.StatusOK, result)

}

func (h *ContentHandler) HandleCreate(c *gin.Context) {
	var createContentDTO DTO.ContentRequestBody

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

	c.JSON(http.StatusCreated, "Conteúdo inserido com sucesso")
}
