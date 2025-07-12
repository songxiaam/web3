package main

import (
	"fmt"
	"log"
	"os"

	"smart-route/pkg/api"
)

func main() {
	port := os.Getenv("SMART_ROUTE_PORT")
	if port == "" {
		port = "8080"
	}

	r := api.NewRouter()
	fmt.Printf("Starting smart-route backend (gin) on :%s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
