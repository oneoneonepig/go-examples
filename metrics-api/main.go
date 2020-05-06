package main

import (
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/oneoneonepig/go-examples/metrics-api/docs"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin"
	"io"
	"os"
)

// @title Simple API for metric collecting
// @version 1.0
// @description

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host node01
// @BasePath /v2
func main() {

	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(apmgin.Middleware(r)) // Middleware for APM Server
	r.Use(cors.Default())       // CORS settings

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/ping", ping)
	r.GET("/metrics", getMetrics)
	r.GET("/error", makeError)
	r.GET("/repair", repair)

	r.GET("/sleep/:duration", sleep)
	r.GET("/connect", connect)
	r.GET("/connect2", connect2)

	r.Run()
}
