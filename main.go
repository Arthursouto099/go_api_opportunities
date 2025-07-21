package main

import (
	"github.com/Arthursouto099/go_api_opportunities/config"
	"github.com/Arthursouto099/go_api_opportunities/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")
	// initialize database config

	// initialize configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization err: %v", err)
		return
	}

	router.Initialize()
}
