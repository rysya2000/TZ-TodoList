package main

import (
	"fmt"
	"log"

	"todolist/config"
	"todolist/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %v\n", err)
	}
	fmt.Println(cfg)

	app.Run(cfg)
}
