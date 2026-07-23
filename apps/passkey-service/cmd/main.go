package main

import (
	"context"
	"log"
	"os"

	"github.com/sentinelvault/sentinelvault/apps/passkey-service/internal/app"
)

func main() {
	ctx := context.Background()
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/passkey-service.yaml"
	}
	if err := app.Run(ctx, configPath); err != nil {
		log.Fatal(err)
	}
}

// maintenance note 63
