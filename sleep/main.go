package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go.elastic.co/apm/module/apmgin"
	"io"
	"os"
)

func main() {

	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(apmgin.Middleware(r)) // Middleware for APM Server
	r.Use(cors.Default())       // CORS settings

	r.GET("/sleep/:duration", sleep)
	r.Run()
}
