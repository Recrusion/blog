package main

import (
	"log"
	"net/http"

	"github.com/Recrusion/blog-api/internal/configs"
	"github.com/Recrusion/blog-api/internal/repository"
)

func main() {
	cfg, err := configs.NewConfig(8080, 5432, "postgres", "postgres", "postgres", "postgres", "localhost")
	if err != nil {
		log.Fatalf("config creation failed: %v", err)
	}

	log.Printf("config created successfully: %+v", cfg)

	db, err := repository.ConnectDatabase(cfg.DatabaseConfig.GetDBDriver(), cfg.DatabaseConfig.GetDSN())
	if err != nil {
		log.Fatalf("failed connection to database: %v", err)
	}
	defer db.Close()
	log.Printf("connection to database successfully: %+v", db)

	http.ListenAndServe(cfg.ServerConfig.GetPort(), nil)
}
