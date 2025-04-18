package handler

import (
	"github.com/JF-Ar/gomarry/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDB()
}
