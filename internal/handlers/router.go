package handlers

import (
	"database/sql"
	"net/http"
	"os"

	"goth/internal/config"
	"goth/internal/gintemplrenderer"
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


 	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{
    FallbackHtmlRenderer: ginHtmlRenderer,
  } 

  r.SetTrustedProxies(nil)

  addStaticRoutes(r)

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

	authMiddleware := m.NewAuthMiddleware(sessionStore, cfg.SessionCookieName)

  r.Use(
  	m.TextHTMLMiddleware(),
  	m.CSPMiddleware(),
  	authMiddleware.AddUserToContext(),
  )

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

func addStaticRoutes(r *gin.Engine) error {

  root, err := os.Getwd()
	if err != nil {
    return err
  }

	// r.Static("/static/*", root)
	r.Static("/static/css", root)
	r.Static("/static/fonts", root)
	r.Static("/static/images", root)
	r.Static("/static/script", root)
	//
  return nil
}
