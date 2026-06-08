package main

import (
	"log"
	"os/exec"
	"remote/internal/config"
	"remote/internal/router"
	"runtime"
)

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	}

	if err != nil {
		log.Printf("failed to open browser: %v", err)
	}
}

func main() {
	log.SetFlags(log.Ltime)

	r := router.Init()

	go openBrowser("http://localhost:" + config.ServerPort + "/qr")

	if err := r.Run(config.ServerAddr); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
