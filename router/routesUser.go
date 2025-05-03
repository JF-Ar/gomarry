package router

import (
	"github.com/JF-Ar/gomarry/handler/userHandler"
	"github.com/JF-Ar/gomarry/repositories/usersRepository"
	"github.com/JF-Ar/gomarry/useCase/userUseCase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userRepo := usersRepository.NewUserRepository(db)
	userUC := userUseCase.NewUserUseCase(userRepo)
	userH := userHandler.NewHandler(userUC)

	uR := rg.Group("/users")
	{
		uR.POST("/register", userH.Create)
	}
}
