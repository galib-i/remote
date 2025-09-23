package controller

import (
	"net/http"
	"remote/internal/service"

	"github.com/gin-gonic/gin"
)

func PressKey(c *gin.Context) {
	key := c.Param("key")

	handleServiceCall(c, func() error {
		return service.PressKey(key)
	}, http.StatusBadRequest)
}
