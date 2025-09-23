package controller

import (
	"net/http"
	"remote/internal/service"

	"github.com/gin-gonic/gin"
)

type CursorMoveRequest struct {
	X float64 `json:"deltaX"`
	Y float64 `json:"deltaY"`
}

func MoveCursor(c *gin.Context) {
	var req CursorMoveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentX, currentY, err := service.GetCursorPos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newX := currentX + req.X
	newY := currentY + req.Y

	if err := service.MoveCursor(newX, newY); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func Click(c *gin.Context) {
	side := c.Param("side")
	leftClick := side == "left"

	if err := service.Click(leftClick); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
