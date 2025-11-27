package loader

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// переменные окружения для конфигурации
type Env struct {
	serverPort int
	dbPort     int
	dbDriver   string
	dbName     string
	dbUsername string
	password   string
	dbHost     string
}

// получения значения переменной окружения по ключу
func getEnv(key string) string {
	return os.Getenv(key)
}

// валидация значения строкового типа
func validateValueString(key, defaultValue string) string {
	value := getEnv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// валидация значения целочисленного типа
func validateValueInt(key string, defaultValue int) int {
	value, _ := strconv.Atoi(getEnv(key))
	if value == 0 {
		return defaultValue
	}
	return value
}

// загрузка переменных окружения и их инициализация в проекте
func LoadFromEnv() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading environment variables: %w", err)
	}

	env := &Env{
		serverPort: validateValueInt("SERVER_PORT", 8080),
		dbPort:     validateValueInt("DB_PORT", 5432),
		dbDriver:   validateValueString("DB_DRIVER", "postgres"),
		dbName:     validateValueString("DB_NAME", "blog"),
		dbUsername: validateValueString("DB_USERNAME", "postgres"),
		password:   validateValueString("PASSWORD", "postgres"),
		dbHost:     validateValueString("DB_HOST", "localhost"),
	}
	return env, nil
}

// геттеры для получения значений переменных окружения вне пакета
func (e *Env) GetServerPort() int {
	return e.serverPort
}

func (e *Env) GetDBPort() int {
	return e.dbPort
}

func (e *Env) GetDBDriver() string {
	return e.dbDriver
}

func (e *Env) GetDBName() string {
	return e.dbName
}

func (e *Env) GetDBUsername() string {
	return e.dbUsername
}

func (e *Env) GetDBPassword() string {
	return e.password
}

func (e *Env) GetDBHost() string {
	return e.dbHost
}
