package interfaces

import (
	"maycms/internal/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ContentHandler struct {
}

func NewContentHandler() *ContentHandler {
	return &ContentHandler{}
}

func (h *ContentHandler) GetContentHandler(c *gin.Context) {

	mockContent := []domain.Content{
		1: {ID: 1, Title: "Meu primeiro conteudo", ContentText: "Lorem ipsum", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		2: {ID: 2, Title: "Meu segundo conteudo", ContentText: "Lorem ipsum", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	// Return the order as JSON
	c.JSON(http.StatusOK, mockContent)
}
