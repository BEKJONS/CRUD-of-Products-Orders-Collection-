package main

import (
	"ulab3/config"
	"ulab3/internal/app"
)

func main() {
	cfg := config.NewConfig()

	app.Run(cfg)
}
