package handlers

import (
	"goth/internal/middleware"
	"goth/internal/store"
	"goth/internal/templates"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct{
  *Handler
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{
    Handler: &Handler{
      Method: http.MethodGet,
      Path: "/",
      Title: "Español con Fabio",
    },
  }
}

func (h *HomeHandler) ServeHTTP(c *gin.Context) {

	_, ok := c.Request.Context().Value(middleware.UserKey).(*store.User)

	if !ok {
		t := templates.Home()
		err := templates.App(t, h.Title).Render(c.Request.Context(), c.Writer)
		if err != nil {
			http.Error(c.Writer, "Error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	t := templates.Home()
	err := templates.App(t, h.Title).Render(c.Request.Context(), c.Writer)
	if err != nil {
		http.Error(c.Writer, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
