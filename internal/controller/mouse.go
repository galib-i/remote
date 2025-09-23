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

	handleServiceCall(c, func() error {
		currentX, currentY, err := service.GetCursorPos()
		if err != nil {
			return err
		}

		newX := currentX + req.X
		newY := currentY + req.Y

		return service.MoveCursor(newX, newY)
	}, http.StatusInternalServerError)
}

func Click(c *gin.Context) {
	side := c.Param("side")
	leftClick := side == "left"

	handleServiceCall(c, func() error {
		return service.Click(leftClick)
	}, http.StatusInternalServerError)
}
