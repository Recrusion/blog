package configs

import (
	"errors"
	"fmt"
	"strconv"
)

// полная конфигурация приложения
type Config struct {
	serverConfig   serverConfig
	databaseConfig databaseConfig
}

// конфигурация сервера
type serverConfig struct {
	port int
}

// создание конфигурации сервера
func newServerConfig(port int) (*serverConfig, error) {
	if port <= 0 || port > 65535 {
		return nil, fmt.Errorf("invalid server port: %d", port)
	}

	return &serverConfig{
		port: port,
	}, nil
}

// конфигурация базы данных
type databaseConfig struct {
	dbDriver string
	dbName   string
	username string
	password string
	host     string
	port     int
}

// создание конфигурации базы данных
func newDatabaseConfig(dbDriver, dbName, username, password, host string, port int) (*databaseConfig, error) {
	if dbDriver == "" {
		return nil, errors.New("database driver cannot be empty")
	}

	if dbName == "" {
		return nil, errors.New("database name cannot be empty")
	}

	if port <= 0 || port > 65535 {
		return nil, fmt.Errorf("invalid database port: %d", port)
	}

	if host == "" {
		return nil, errors.New("database host cannot be empty")
	}

	return &databaseConfig{
		dbDriver: dbDriver,
		dbName:   dbName,
		username: username,
		password: password,
		host:     host,
		port:     port,
	}, nil
}

// создание полной конфигурации приложения
func NewConfig(serverPort, dbPort int, dbDriver, dbName, username, password, host string) (*Config, error) {
	srvConfig, err := newServerConfig(serverPort)
	if err != nil {
		return nil, err
	}

	dbConfig, err := newDatabaseConfig(dbDriver, dbName, username, password, host, dbPort)
	if err != nil {
		return nil, err
	}

	return &Config{
		serverConfig:   *srvConfig,
		databaseConfig: *dbConfig,
	}, nil
}

// геттеры для получения данных конфигурации вне пакета
func (c *Config) GetServerConfig() *serverConfig {
	return &c.serverConfig
}

func (c *Config) GetDatabaseConfig() *databaseConfig {
	return &c.databaseConfig
}

func (s *serverConfig) GetPort() string {
	return ":" + strconv.Itoa(s.port)
}

func (d *databaseConfig) GetDBDriver() string {
	return d.dbDriver
}

func (d *databaseConfig) GetDSN() string {
	switch d.dbDriver {
	case "postgres":
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			d.host, d.port, d.username, d.password, d.dbName)
	default:
		return ""
	}
}
