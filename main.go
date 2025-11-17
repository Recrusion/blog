package main

import (
	"log"
	"net/http"

	"github.com/Recrusion/blog-api/internal/configs"
	"github.com/Recrusion/blog-api/internal/loader"
	"github.com/Recrusion/blog-api/internal/repository"
)

func main() {
	env, err := loader.LoadFromEnv()
	if err != nil {
		log.Printf(".env undefined: %v", err)
	}

	cfg, err := configs.NewConfig(env.GetServerPort(), env.GetDBPort(), env.GetDBDriver(), env.GetDBName(), env.GetDBUsername(), env.GetDBPassword(), env.GetDBHost())
	if err != nil {
		log.Fatalf("config creation failed: %v", err)
	}

	log.Printf("config created successfully")

	db, err := repository.ConnectDatabase(cfg.GetDatabaseConfig().GetDBDriver(), cfg.GetDatabaseConfig().GetDSN())
	if err != nil {
		log.Fatalf("failed connection to database: %v", err)
	}
	defer db.Close()
	log.Printf("connection to database successfully")

	if err := http.ListenAndServe(cfg.GetServerConfig().GetPort(), nil); err != nil {
		log.Fatalf("server is down: %v", err)
	}
}
