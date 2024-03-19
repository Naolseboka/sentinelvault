package main

import (
	"context"
	"log"
	"os"

	"github.com/sentinelvault/sentinelvault/apps/search-service/internal/app"
)

func main() {
	ctx := context.Background()
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/search-service.yaml"
	}
	if err := app.Run(ctx, configPath); err != nil {
		log.Fatal(err)
	}
}
