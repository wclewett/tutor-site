package handlers

import (
	"context"
	"goth/internal/gintemplrenderer"
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
    ctx := context.WithValue(
      context.Background(), 
      middleware.NonceKey, 
      c.Request.Context().Value(middleware.NonceKey),
    )
    // fmt.Println(ctx.Value(middleware.NonceKey))
    r := gintemplrenderer.New(
      ctx, 
      http.StatusOK, 
      templates.App(templates.Home(), h.Title),
    )
    // .Render(c.Request.Context(), c.Writer)
    c.Render(http.StatusOK, r)
		// if err != nil {
		// 	http.Error(c.Writer, "Error rendering template", http.StatusInternalServerError)
		// 	return
		// }
		return
	}

    r := gintemplrenderer.New(
      c.Request.Context(), 
      http.StatusOK, 
      templates.App(templates.Home(), h.Title),
    )
    c.Render(http.StatusOK, r)
	// t := templates.Home()
	// err := templates.App(t, h.Title).Render(c.Request.Context(), c.Writer)
	// if err != nil {
	// 	http.Error(c.Writer, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}
