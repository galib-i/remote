package main

import (
	"os/exec"
	"runtime"

	"remote/internal/router"
)

func openBrowser(url string) {
	switch runtime.GOOS {
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	}
}

func main() {
	r := router.Init()
	go openBrowser("http://localhost:8080")
	r.Run(":8080")
}
