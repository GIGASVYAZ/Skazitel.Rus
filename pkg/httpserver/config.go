package httpserver

import (
	"os"
	"strconv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	URL      string
	MaxConns int32
	MinConns int32
}

type ServerConfig struct {
	Port string
}

func New() Config {
	return Config{
		Database: DatabaseConfig{
			URL:      getEnv("DATABASE_URL", "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"),
			MaxConns: int32(getEnvAsInt("DB_MAX_CONNS", 25)),
			MinConns: int32(getEnvAsInt("DB_MIN_CONNS", 5)),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valStr := getEnv(name, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}
