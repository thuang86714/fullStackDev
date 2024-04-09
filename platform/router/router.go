package router

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/thuang86714/fullStackDev/platform/authenticator"
	"github.com/thuang86714/fullStackDev/platform/middleware"
	"github.com/thuang86714/fullStackDev/web/app/callback"
	"github.com/thuang86714/fullStackDev/web/app/login"
	"github.com/thuang86714/fullStackDev/web/app/logout"
	"github.com/thuang86714/fullStackDev/web/app/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated,user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}