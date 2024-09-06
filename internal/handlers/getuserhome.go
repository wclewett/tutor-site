package handlers

import (
	"goth/internal/templates"
	"net/http"
)

type GetUserHomeHandler struct{}

func NewGetUserHomeHandler() *GetUserHomeHandler {
	return &GetUserHomeHandler{}
}

func (h *GetUserHomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.UserHome()
	err := templates.App(c, "Inicio de usario").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}
