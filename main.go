package main

import (
	"github.com/JF-Ar/gomarry/config"
	r "github.com/JF-Ar/gomarry/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	if er := config.Init(); er != nil {
		logger.ErrorF("Error initializing config: %v", er)
	}

	r.Router()
}
