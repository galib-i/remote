package controller

import (
	"net/http"

	"remote/internal/service"

	"github.com/gin-gonic/gin"
)

func VolumeUp(c *gin.Context) {
	err := service.VolumeUp()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Volume increased successfully",
	})
}

func VolumeDown(c *gin.Context) {
	err := service.VolumeDown()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Volume decreased successfully",
	})
}

func MuteVolume(c *gin.Context) {
	err := service.MuteVolume()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Volume muted successfully",
	})
}

func UnmuteVolume(c *gin.Context) {
	err := service.UnmuteVolume()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Volume unmuted successfully",
	})
}
