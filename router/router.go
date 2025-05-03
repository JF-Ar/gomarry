package router

import (
	"github.com/JF-Ar/gomarry/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	handler.InitHandler()

	router := gin.Default()
	api := router.Group("/api/v1")

	{
		api.GET("/", handler.PingHandler)
		RegisterUserRoutes(api, db)
	}

	return router

}
