package configs

import (
	"errors"
	"fmt"
	"strconv"
)

type Config struct {
	ServerConfig   ServerConfig
	DatabaseConfig DatabaseConfig
}

type ServerConfig struct {
	port int
}

func newServerConfig(port int) (*ServerConfig, error) {
	if port <= 0 || port > 65535 {
		return nil, fmt.Errorf("invalid server port: %d", port)
	}

	return &ServerConfig{
		port: port,
	}, nil
}

type DatabaseConfig struct {
	dbDriver string
	dbName   string
	username string
	password string
	host     string
	port     int
}

func newDatabaseConfig(dbDriver, dbName, username, password, host string, port int) (*DatabaseConfig, error) {
	if dbDriver == "" {
		return nil, errors.New("database driver cannot be empty")
	}

	if dbName == "" {
		return nil, errors.New("database name cannot be empty")
	}

	if port <= 0 || port > 65535 {
		return nil, fmt.Errorf("invalid database port: %d", port)
	}

	return &DatabaseConfig{
		dbDriver: dbDriver,
		dbName:   dbName,
		username: username,
		password: password,
		host:     host,
		port:     port,
	}, nil
}

func NewConfig(serverPort, dbPort int, dbDriver, dbName, username, password, host string) (*Config, error) {
	serverConfig, err := newServerConfig(serverPort)
	if err != nil {
		return nil, err
	}

	databaseConfig, err := newDatabaseConfig(dbDriver, dbName, username, password, host, dbPort)
	if err != nil {
		return nil, err
	}

	return &Config{
		ServerConfig:   *serverConfig,
		DatabaseConfig: *databaseConfig,
	}, nil
}

func (s *ServerConfig) GetPort() string {
	return ":" + strconv.Itoa(s.port)
}

func (d *DatabaseConfig) GetDSN() string {
	switch d.dbDriver {
	case "postgres":
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			d.host, d.port, d.username, d.password, d.dbName)
	default:
		return ""
	}
}
