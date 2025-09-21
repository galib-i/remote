package main

import (
	"remote/internal/router"
)

func main() {
	r := router.Init()
	r.Run("0.0.0.0:12345")
}
