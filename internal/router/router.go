package router

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	r := gin.Default()

	r.Static("/", "./web")

	return r
}
