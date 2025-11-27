package main

import (
	"log"

	"github.com/Recrusion/blog-api/internal/configs"
	"github.com/Recrusion/blog-api/internal/loader"
	"github.com/Recrusion/blog-api/internal/repository"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// инициализация переменных окружения
	env, err := loader.LoadFromEnv()
	if err != nil {
		log.Fatalf("[ERROR] .env undefined: %v", err)
	}

	// инициализация конфига
	cfg, err := configs.NewConfig(env.GetServerPort(), env.GetDBPort(), env.GetDBDriver(), env.GetDBName(), env.GetDBUsername(), env.GetDBPassword(), env.GetDBHost())
	if err != nil {
		log.Fatalf("[ERROR] config creation failed: %v", err)
	}

	log.Printf("[INFO] config created successfully")

	// создание объекта базы данных
	db, err := repository.ConnectDatabase(cfg.GetDatabaseConfig().GetDBDriver(), cfg.GetDatabaseConfig().GetDSN())
	if err != nil {
		log.Fatalf("[ERROR] failed connection to database: %v", err)
	}
	defer db.Close()
	log.Printf("[INFO] connection to database successfully")

	// старт сервера
	if err = e.Start(cfg.GetServerConfig().GetPort()); err != nil {
		log.Fatalf("[ERROR] failed to start server: %v", err)
	}
}
