package main

import (
	"log"
	"net/http"

	"github.com/Recrusion/blog-api/internal/configs"
)

func main() {
	cfg, err := configs.NewConfig(8080, 5432, "postgres", "postgres", "postgres", "postgres", "localhost")
	if err != nil {
		log.Fatalf("config creation failed: %v", err)
	}

	log.Printf("config created successfully: %+v", cfg)

	http.ListenAndServe(cfg.ServerConfig.GetPort(), nil)
}
