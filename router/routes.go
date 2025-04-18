package router

import (
	"github.com/JF-Ar/gomarry/handler"
	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {

	handler.InitHandler()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/", handler.PingHandler)
	}
}
