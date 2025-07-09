package main

import (
	"log"

	"github.com/Luiz0311/estoque-go/config"
	"github.com/Luiz0311/estoque-go/router"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	router.Initialize()
}
