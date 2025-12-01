package main

import (
	"log/slog"
	"os"

	"github.com/Recrusion/blog-api/internal/configs"
	"github.com/Recrusion/blog-api/internal/handlers"
	"github.com/Recrusion/blog-api/internal/loader"
	"github.com/Recrusion/blog-api/internal/repository"
	"github.com/Recrusion/blog-api/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// инициализация логгера
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// инициализация переменных окружения
	env, err := loader.LoadFromEnv()
	if err != nil {
		logger.Error("error loading env variables", err)
	}

	// инициализация конфига
	cfg, err := configs.NewConfig(env.GetServerPort(), env.GetDBPort(), env.GetDBDriver(), env.GetDBName(), env.GetDBUsername(), env.GetDBPassword(), env.GetDBHost())
	if err != nil {
		logger.Error("config creation failed", err)
	}

	logger.Info("config created successfully")

	// создание объекта базы данных
	db, err := repository.ConnectDatabase(cfg.GetDatabaseConfig().GetDBDriver(), cfg.GetDatabaseConfig().GetDSN())
	if err != nil {
		logger.Error("failed connection to database", err)
	}
	defer db.Close()
	logger.Info("connection to database successfully")

	database := repository.NewDatabase(db)
	service := service.NewService(database)
	handler := handlers.NewHandlers(service)

	e.POST("/createpost", handler.CreatePost)
	e.GET("/getpost:id", handler.GetPost)
	// старт сервера
	if err = e.Start(cfg.GetServerConfig().GetPort()); err != nil {
		logger.Error("failed to start server", err)
	}
}
