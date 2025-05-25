package handlers

import (
	"goth/internal/templates"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserHomeHandler struct{}

func NewGetUserHomeHandler() *GetUserHomeHandler {
	return &GetUserHomeHandler{}
}

func (h *GetUserHomeHandler) ServeHTTP(c *gin.Context) {
	t := templates.UserHome()
	err := templates.App(t, "Inicio de usario").Render(c.Request.Context(), c.Writer)

	if err != nil {
		http.Error(c.Writer, "Error rendering template", http.StatusInternalServerError)
		return
	}

}
