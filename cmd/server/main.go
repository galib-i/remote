package main

import (
	"os/exec"
	"remote/internal/config"
	"remote/internal/router"
	"runtime"
)

func openBrowser(url string) {
	switch runtime.GOOS {
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "linux":
		exec.Command("xdg-open", url).Start()
	}
}

func main() {
	r := router.Init()

	go openBrowser("http://localhost:" + config.ServerPort + "/qr")

	r.Run(config.ServerAddr)
}
