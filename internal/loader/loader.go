package loader

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	serverPort string
	dbPort     string
	dbDriver   string
	dbName     string
	username   string
	password   string
	dbHost     string
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func validateValue(key string, defaultValue string) string {
	value := getEnv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func LoadFromEnv() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading environment variables: %w", err)
	}

	env := &Env{
		serverPort: validateValue("SERVER_PORT", "8080"),
		dbPort:     validateValue("DB_PORT", "5432"),
		dbDriver:   validateValue("DB_DRIVER", "postgres"),
		dbName:     validateValue("DB_NAME", "postgres"),
		username:   validateValue("USERNAME", "postgres"),
		password:   validateValue("PASSWORD", "postgres"),
		dbHost:     validateValue("DB_HOST", "localhost"),
	}
	return env, nil
}

func (e *Env) GetServerPort() string {
	return e.serverPort
}

func (e *Env) GetDBPort() string {
	return e.dbPort
}

func (e *Env) GetDBDriver() string {
	return e.dbDriver
}

func (e *Env) GetDBName() string {
	return e.dbName
}

func (e *Env) GetDBUsername() string {
	return e.username
}

func (e *Env) GetDBPassword() string {
	return e.password
}

func (e *Env) GetDBHost() string {
	return e.dbHost
}
