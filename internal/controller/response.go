package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleServiceCall(c *gin.Context, serviceFunc func() error, errorStatus int) {
	if err := serviceFunc(); err != nil {
		c.JSON(errorStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
