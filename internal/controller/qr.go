package controller

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func ShowQR(c *gin.Context) {
	ip, err := getLocalIP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get local IP"})
		return
	}

	url := fmt.Sprintf("http://%s:12345", ip)

	// Generate QR code as PNG bytes
	qrCode, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Header("Cache-Control", "no-cache")
	c.Data(http.StatusOK, "image/png", qrCode)
}
