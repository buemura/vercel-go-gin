package api

import (
	"net/http"

	"github.com/buemura/vercel-go-gin/handler"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func setupRouter(r *gin.RouterGroup) {
	r.GET("/ping", handler.Ping)
	r.GET("/hello", handler.Hello)
}

func init() {
	app = gin.New()
	app.NoRoute(handler.ErrRouter)

	router := app.Group("/api")
	setupRouter(router)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
