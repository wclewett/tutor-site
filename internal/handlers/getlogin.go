package handlers

import (
	"goth/internal/templates"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetLoginHandler struct{}

func NewGetLoginHandler() *GetLoginHandler {
	return &GetLoginHandler{}
}

func (h *GetLoginHandler) ServeHTTP(c *gin.Context) {
	t := templates.Login()
	err := templates.App(t, "Acceder").Render(c.Request.Context(), c.Writer)

	if err != nil {
		http.Error(c.Writer, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
