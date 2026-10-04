package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisHost        string
	RedisPort        string
	RedisPassword    string
	RedisDB          string
	Port             string
	DatabaseURL      string
	JwtSecret        string
	EmailWorkerCount string
}

func Load() (Config, error) {

	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	config := Config{
		RedisHost:        os.Getenv("REDIS_HOST"),
		RedisPort:        os.Getenv("REDIS_PORT"),
		RedisPassword:    os.Getenv("REDIS_PASSWORD"),
		RedisDB:          os.Getenv("REDIS_DB"),
		Port:             os.Getenv("PORT"),
		DatabaseURL:      os.Getenv("DATABASE_URL"),
		JwtSecret:        os.Getenv("JWT_SECRET"),
		EmailWorkerCount: os.Getenv("EMAIL_WORKER_COUNT"),
	}

	if config.RedisHost == "" {
		return Config{}, errors.New("RedisHost is required")
	} else if config.RedisPort == "" {
		return Config{}, errors.New("RedisPort is required")
	} else if config.Port == "" {
		return Config{}, errors.New("Port is required")
	} else if config.DatabaseURL == "" {
		return Config{}, errors.New("MongoURI is required")
	} else if config.JwtSecret == "" {
		return Config{}, errors.New("JWT_SECRET is required")
	} else if config.EmailWorkerCount == "" {
		return Config{}, errors.New("EMAIL_WORKER_COUNT is required")
	}

	return config, nil
}
