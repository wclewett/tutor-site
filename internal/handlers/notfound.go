package handlers

import (
	"goth/internal/templates"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotFoundHandler struct{}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}

func (h *NotFoundHandler) ServeHTTP(c *gin.Context) {
	t := templates.NotFound()
	err := templates.App(t, "Not Found").Render(c.Request.Context(), c.Writer)

	if err != nil {
		http.Error(c.Writer, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
