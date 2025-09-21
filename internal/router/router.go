package router

import (
	"remote/internal/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.Static("/", "./web")

	r.POST("/volume-up", controller.VolumeUp)
	r.POST("/volume-down", controller.VolumeDown)
	r.POST("/mute-volume", controller.MuteVolume)
	r.POST("/unmute-volume", controller.UnmuteVolume)
	r.POST("/move-mouse", controller.MoveMouse)

	return r
}
