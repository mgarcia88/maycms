package api

import (
	DTO "maycms/Internal/Adapters/Driving/Api/DTOs"
	"maycms/internal/domain/application"
	"maycms/internal/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContentHandler struct {
	service application.ContentService
}

func NewContentHandler(s application.ContentService) *ContentHandler {
	return &ContentHandler{service: s}
}

func (h *ContentHandler) GetContentHandler(c *gin.Context) {
	contents := h.service.GetAllContents()
	c.JSON(http.StatusOK, contents)
}

func (h *ContentHandler) GetContentByIDHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

	result := h.service.GetContentById(id)
	c.JSON(http.StatusOK, result)

}

func (h *ContentHandler) CreateContentHandler(c *gin.Context) {
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

	err = h.service.CreateContent(content)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Conteúdo inserido com sucesso")
}
