package router

import (
	"remote/internal/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.Static("/", "./web")

	r.POST("/volume/:direction", controller.AdjustVolume) // :direction -> "up"/"down"
	r.POST("/toggle-mute", controller.ToggleMute)
	r.POST("/move-cursor", controller.MoveCursor)
	r.POST("/click/:side", controller.Click) // :side -> "left"/"right"
	r.POST("/press-key/:key", controller.PressKey)

	return r
}
