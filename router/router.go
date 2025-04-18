package router

import "github.com/gin-gonic/gin"

func Router() {
	router := gin.Default()

	registerRoutes(router)
	err := router.Run(":8080")
	if err != nil {
		return
	}
	
}
