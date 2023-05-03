package main

import (
	"context"
	"log"
	"os"

	"github.com/sentinelvault/sentinelvault/apps/attachment-service/internal/app"
)

func main() {
	ctx := context.Background()
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/attachment-service.yaml"
	}
	if err := app.Run(ctx, configPath); err != nil {
		log.Fatal(err)
	}
}
