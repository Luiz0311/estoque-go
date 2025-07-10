package main

import (
	"github.com/Luiz0311/estoque-go/config"
	"github.com/Luiz0311/estoque-go/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Err("erro ao carregar configuração", err)
	}

	router.Initialize()
}
