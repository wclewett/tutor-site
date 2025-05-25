package handlers

import (
	b64 "encoding/base64"
	"fmt"
	"goth/internal/hash"
	"goth/internal/store"
	"goth/internal/templates"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PostLoginHandler struct {
	userStore         store.UserStore
	sessionStore      store.SessionStore
	passwordhash      hash.PasswordHash
	sessionCookieName string
}

type PostLoginHandlerParams struct {
	UserStore         store.UserStore
	SessionStore      store.SessionStore
	PasswordHash      hash.PasswordHash
	SessionCookieName string
}

func NewPostLoginHandler(params PostLoginHandlerParams) *PostLoginHandler {
	return &PostLoginHandler{
		userStore:         params.UserStore,
		sessionStore:      params.SessionStore,
		passwordhash:      params.PasswordHash,
		sessionCookieName: params.SessionCookieName,
	}
}

func (h *PostLoginHandler) ServeHTTP(c *gin.Context) {

	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	user, err := h.userStore.GetUser(email)

	if err != nil {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		t := templates.LoginError()
		t.Render(c.Request.Context(), c.Writer)
		return
	}

	passwordIsValid, err := h.passwordhash.ComparePasswordAndHash(password, user.Password)

	if err != nil || !passwordIsValid {
	  c.Writer.WriteHeader(http.StatusUnauthorized)
		t := templates.LoginError()
		t.Render(c.Request.Context(), c.Writer)
		return
	}

	session, err := h.sessionStore.CreateSession(&store.Session{
		UserID: user.ID,
	})

	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID := user.ID
	sessionID := session.SessionID

	cookieValue := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%d", sessionID, userID)))

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:     h.sessionCookieName,
		Value:    cookieValue,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(c.Writer, &cookie)

	c.Writer.Header().Set("HX-Redirect", "/")
	c.Writer.WriteHeader(http.StatusOK)
}
