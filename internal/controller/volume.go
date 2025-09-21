package controller

import (
	"net/http"

	"remote/internal/service"

	"github.com/gin-gonic/gin"
)

func AdjustVolume(c *gin.Context) {
	direction := c.Param("direction")
	increase := direction == "up"

	if err := service.AdjustVolume(increase); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func ToggleMute(c *gin.Context) {
	if err := service.ToggleMute(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
