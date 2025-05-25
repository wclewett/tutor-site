package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PostLogoutHandler struct {
	sessionCookieName string
}

type PostLogoutHandlerParams struct {
	SessionCookieName string
}

func NewPostLogoutHandler(params PostLogoutHandlerParams) *PostLogoutHandler {
	return &PostLogoutHandler{
		sessionCookieName: params.SessionCookieName,
	}
}

func (h *PostLogoutHandler) ServeHTTP(c *gin.Context) {

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    h.sessionCookieName,
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
		Path:    "/",
	})

	http.Redirect(c.Writer, c.Request, "/", http.StatusSeeOther)
}
