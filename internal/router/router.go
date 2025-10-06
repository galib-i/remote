package router

import (
	"remote/internal/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/qr", controller.ShowQR)
		api.POST("/volume/:direction", controller.AdjustVolume) // :direction -> "up"/"down"
		api.POST("/toggle-mute", controller.ToggleMute)
		api.POST("/move-cursor", controller.MoveCursor)
		api.POST("/click/:side", controller.Click) // :side -> "left"/"right"
		api.POST("/press-key", controller.PressKey)
	}

	r.StaticFile("/", "./web/index.html")
	r.StaticFile("/qr", "./web/qr.html")

	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")

	return r
}
