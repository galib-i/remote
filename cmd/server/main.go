package main

import (
	"os/exec"
	"remote/internal/router"
	"runtime"
)

func openBrowser(url string) {
	switch runtime.GOOS {
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	}
}

func main() {
	r := router.Init()
	go openBrowser("http://localhost:12345/qr")
	r.Run("0.0.0.0:12345")
}
