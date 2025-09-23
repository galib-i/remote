package controller

import (
	"net/http"

	"remote/internal/service"

	"github.com/gin-gonic/gin"
)

func AdjustVolume(c *gin.Context) {
	direction := c.Param("direction")
	increase := direction == "up"

	handleServiceCall(c, func() error {
		return service.AdjustVolume(increase)
	}, http.StatusInternalServerError)
}

func ToggleMute(c *gin.Context) {
	handleServiceCall(c, func() error {
		return service.ToggleMute()
	}, http.StatusInternalServerError)
}
