package api

import (
	"maycms/internal/adapters/driven/infra/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContentHandler struct {
}

func NewContentHandler() *ContentHandler {
	return &ContentHandler{}
}

func (h *ContentHandler) GetContentHandler(c *gin.Context) {
	contentRepository := repositories.NewContentRepository()

	mockContent := contentRepository.GetAllContents()

	// Return the order as JSON
	c.JSON(http.StatusOK, mockContent)
}

func (h *ContentHandler) GetContentByIDHandler(c *gin.Context) {
	contentRepository := repositories.NewContentRepository()
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

	mockContent := contentRepository.GetContentById(id)

	// Return the order as JSON
	c.JSON(http.StatusOK, mockContent)
}
