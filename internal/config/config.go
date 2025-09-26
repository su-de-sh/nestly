package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/su-de-sh/nestly/utils"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Name     string
	Port     string
	SSLMode  string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	config := &Config{
		Server: ServerConfig{
			Port: utils.GetEnv("PORT", "5001"),
		},
		Database: DatabaseConfig{
			Host:     utils.GetEnv("DB_HOST", "localhost"),
			Port:     utils.GetEnv("DB_PORT", "5432"),
			User:     utils.GetEnv("DB_USER", "postgres"),
			Password: utils.GetEnv("DB_PASSWORD", ""),
			Name:     utils.GetEnv("DB_NAME", "nestly"),
			SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
		},
	}

	return config, nil
}

func (c *Config) GetDatabaseConnectionString() string {
	return "host=" + c.Database.Host +
		" port=" + c.Database.Port +
		" user=" + c.Database.User +
		" password=" + c.Database.Password +
		" dbname=" + c.Database.Name +
		" sslmode=" + c.Database.SSLMode
}
