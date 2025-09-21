package controller

import (
	"fmt"
	"net/http"
	"remote/internal/service"

	"github.com/gin-gonic/gin"
)

type MouseMoveRequest struct {
	X float64 `json:"deltaX"`
	Y float64 `json:"deltaY"`
}

func MoveMouse(c *gin.Context) {
	var req MouseMoveRequest
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

	err = service.MoveMouse(newX, newY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mouse moved"})
	fmt.Println(newX, newY)
}
