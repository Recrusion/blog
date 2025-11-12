package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Recrusion/blog-api/internal/configs"
	"github.com/Recrusion/blog-api/internal/loader"
	"github.com/Recrusion/blog-api/internal/repository"
)

func main() {
	env, err := loader.LoadFromEnv()
	if err != nil {
		log.Printf("error: %v", err)
	}

	serverPort, _ := strconv.Atoi(env.GetServerPort())
	dbPort, _ := strconv.Atoi(env.GetDBPort())
	cfg, err := configs.NewConfig(serverPort, dbPort, env.GetDBDriver(), env.GetDBName(), env.GetDBUsername(), env.GetDBPassword(), env.GetDBHost())
	if err != nil {
		log.Fatalf("config creation failed: %v", err)
	}

	log.Printf("config created successfully: %+v", cfg)

	db, err := repository.ConnectDatabase(cfg.GetDatabaseConfig().GetDBDriver(), cfg.GetDatabaseConfig().GetDSN())
	if err != nil {
		log.Fatalf("failed connection to database: %v", err)
	}
	defer db.Close()
	log.Printf("connection to database successfully: %+v", db)

	http.ListenAndServe(cfg.GetServerConfig().GetPort(), nil)
}
