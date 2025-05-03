package main

import (
	"github.com/JF-Ar/gomarry/config"
	"github.com/JF-Ar/gomarry/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	db, err := config.Init()
	if err != nil {
		logger.ErrorF("Error initializing config: %v", err)
	}

	r := router.SetupRouter(db)
	if err := r.Run(":8080"); err != nil {
		logger.ErrorF("Error initializing server: %v", err)
		panic(err)
	}
}
