package handlers

import (
	"goth/internal/store"
	"goth/internal/templates"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostRegisterHandler struct {
	userStore store.UserStore
}

type PostRegisterHandlerParams struct {
	UserStore store.UserStore
}

func NewPostRegisterHandler(params PostRegisterHandlerParams) *PostRegisterHandler {
	return &PostRegisterHandler{
		userStore: params.UserStore,
	}
}

func (h *PostRegisterHandler) ServeHTTP(c *gin.Context) {
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	err := h.userStore.CreateUser(email, password)

	if err != nil {

		c.Writer.WriteHeader(http.StatusBadRequest)
		t := templates.RegisterError()
		t.Render(c.Request.Context(), c.Writer)
		return
	}

	t := templates.RegisterSuccess()
	err = t.Render(c.Request.Context(), c.Writer)

	if err != nil {
		http.Error(c.Writer, "error rendering template", http.StatusInternalServerError)
		return
	}

}
