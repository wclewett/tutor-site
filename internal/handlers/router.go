package handlers

import (
	"database/sql"
	"net/http"
	"os"

	"goth/internal/config"
	"goth/internal/hash/passwordhash"
	m "goth/internal/middleware"
	"goth/internal/store/dbstore"

	"github.com/gin-gonic/gin"
)

const (
  fs_path = "/static"
)

func Router(cfg *config.Config, db *sql.DB) (http.Handler, error) {
  r := gin.Default()
  
  root, err := os.Getwd()
	if err != nil {
    return r, err
  }
	
  passwordhash := passwordhash.NewHPasswordHash()

	userStore := dbstore.NewUserStore(
		dbstore.NewUserStoreParams{
			DB:           db,
			PasswordHash: passwordhash,
		},
	)

	sessionStore := dbstore.NewSessionStore(
		dbstore.NewSessionStoreParams{
			DB: db,
		},
	)

  fileServer := http.FileServer(http.Dir(root+fs_path))

  // todo: fix file handler
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	authMiddleware := m.NewAuthMiddleware(sessionStore, cfg.SessionCookieName)

  // r.Use(
  // 	m.TextHTMLMiddleware,
  // 	m.CSPMiddleware,
  // 	authMiddleware.AddUserToContext,
  // )

  r.NoRoute(NewNotFoundHandler().ServeHTTP)

  r.GET("/", NewHomeHandler().ServeHTTP)

  r.GET("/crear", NewGetRegisterHandler().ServeHTTP)
  r.POST("/crear", NewPostRegisterHandler(PostRegisterHandlerParams{
    UserStore: userStore,
  }).ServeHTTP)

  r.GET("/acceder", NewGetLoginHandler().ServeHTTP)
  r.POST("/acceder", NewPostLoginHandler(PostLoginHandlerParams{
    UserStore:         userStore,
    SessionStore:      sessionStore,
    PasswordHash:      passwordhash,
    SessionCookieName: cfg.SessionCookieName,
  }).ServeHTTP)

  // r.Get("/reservar", NewGetReserveHandler().ServeHTTP)
  // r.Post("/reservar", NewPostReserveHandler(PostReserveHandlerParams{
  // 	UserStore:         userStore,
  // 	SessionStore:      sessionStore,
  // 	PasswordHash:      passwordhash,
  // 	SessionCookieName: cfg.SessionCookieName,
  // }).ServeHTTP)

  r.GET("/iniciodeusario", NewGetUserHomeHandler().ServeHTTP)


  r.POST("/salir", NewPostLogoutHandler(PostLogoutHandlerParams{
    SessionCookieName: cfg.SessionCookieName,
  }).ServeHTTP)
  return r, nil
}


func addRoutes(r *gin.Engine) {
  // add all routes in switch
}
