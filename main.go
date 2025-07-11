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
		logger.Errf("erro ao carregar configuração da base de dados: %v", err)
		return
	}

	router.Initialize()
}
