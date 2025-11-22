package main

import (
	"log"
	"net/http"

	"github.com/Recrusion/blog-api/internal/configs"
	"github.com/Recrusion/blog-api/internal/loader"
	"github.com/Recrusion/blog-api/internal/repository"
)

func main() {
	// Инициализация переменных окружения
	env, err := loader.LoadFromEnv()
	if err != nil {
		log.Fatalf("[ERROR] .env undefined: %v", err)
	}

	// Инициализация конфига
	cfg, err := configs.NewConfig(env.GetServerPort(), env.GetDBPort(), env.GetDBDriver(), env.GetDBName(), env.GetDBUsername(), env.GetDBPassword(), env.GetDBHost())
	if err != nil {
		log.Fatalf("[ERROR] config creation failed: %v", err)
	}

	log.Printf("[INFO] config created successfully")

	// Создание объекта базы данных
	db, err := repository.ConnectDatabase(cfg.GetDatabaseConfig().GetDBDriver(), cfg.GetDatabaseConfig().GetDSN())
	if err != nil {
		log.Fatalf("[ERROR] failed connection to database: %v", err)
	}
	defer db.Close()
	log.Printf("[INFO] connection to database successfully")

	// Инициализация сервера
	if err = http.ListenAndServe(cfg.GetServerConfig().GetPort(), nil); err != nil {
		log.Fatalf("[ERROR] server is down: %v", err)
	}
}
