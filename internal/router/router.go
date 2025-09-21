package router

import (
	"remote/internal/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.Static("/", "./web")

	r.POST("/volume/:direction", controller.AdjustVolume) // :direction will be "up" or "down"
	r.POST("/toggle-mute", controller.ToggleMute)
	r.POST("/move-mouse", controller.MoveCursor)

	return r
}
