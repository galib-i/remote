package controller

import (
	"net/http"
	"remote/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func ShowQR(c *gin.Context) {
	if config.ServerURL == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to determine server URL"})
		return
	}

	// Generate QR code as PNG bytes
	qrCode, err := qrcode.Encode(config.ServerURL, qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Header("Cache-Control", "no-cache")
	c.Data(http.StatusOK, "image/png", qrCode)
}
