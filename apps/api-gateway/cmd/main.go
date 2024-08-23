package main

import (
	"context"
	"log"
	"os"

	"github.com/sentinelvault/sentinelvault/apps/api-gateway/internal/gateway"
)

func main() {
	ctx := context.Background()
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/api-gateway.yaml"
	}
	if err := gateway.Run(ctx, configPath); err != nil {
		log.Fatal(err)
	}
}
