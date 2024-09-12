package config

import (
    "log"
    "os"
)

type Config struct {
    DatabaseURL string
}

func LoadConfig() Config {
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL muhim o'zgaruvchi topilmadi")
    }
    return Config{
        DatabaseURL: dbURL,
    }
}
