package router

import (
	"remote/internal/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		api.GET("/qr", controller.ShowQR)
		api.GET("/ws", controller.Ws)
	}

	r.StaticFile("/", "./web/index.html")
	r.StaticFile("/qr", "./web/qr.html")

	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")

	return r
}
