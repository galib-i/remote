package config

import (
	"fmt"
	"log"
	"math/rand"
	"net"
)

var (
	ServerPort string
	ServerAddr string
	ServerURL  string
)

func init() {
	ServerPort = fmt.Sprintf("%d", rand.Intn(10000)+50000) // Random port between 50000-59999
	ServerAddr = "0.0.0.0:" + ServerPort

	if ip, err := getLocalIP(); err == nil {
		ServerURL = fmt.Sprintf("http://%s:%s", ip, ServerPort)
	} else {
		log.Printf("failed to determine local IP: %v", err)
	}
}

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}
