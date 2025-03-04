package api

import (
	"maycms/internal/domain/application"
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
