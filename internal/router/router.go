package router

import (
	"io/fs"
	"net/http"
	"remote/internal/controller"

	"github.com/gin-gonic/gin"
)

func Init(frontendFS fs.FS) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	webFS, _ := fs.Sub(frontendFS, "web")
	cssFS, _ := fs.Sub(frontendFS, "web/css")
	jsFS, _ := fs.Sub(frontendFS, "web/js")

	indexHTML, err := fs.ReadFile(webFS, "index.html")
	if err != nil {
		panic("failed to read embedded index.html: " + err.Error())
	}

	qrHTML, err := fs.ReadFile(webFS, "qr.html")
	if err != nil {
		panic("failed to read embedded qr.html: " + err.Error())
	}

	api := r.Group("/api")
	{
		api.GET("/qr", controller.ShowQR)
		api.GET("/ws", controller.Ws)
	}

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHTML)
	})
	r.GET("/qr", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", qrHTML)
	})

	r.StaticFS("/css", http.FS(cssFS))
	r.StaticFS("/js", http.FS(jsFS))

	return r
}
