package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"log"
)

type Config struct {
	Env 	string
	Port 	string
	DB 		DBConfig

}

type DBConfig struct {
	Host	string
	Port 	int
	User 	string
	Pass 	string
	Name	string
}

func Load() *Config {
	_ = godotenv.Load()

	dbPort, err := strconv.Atoi(getEnv("DB_PORT","5432"))
	if err != nil {
		log.Fatalf("invalid DB_PORT: %v",err)
	}

	return &Config{
		Env: getEnv("ENV", "DEV"),
		Port: getEnv("APP_PORT", "8080"),
		DB: DBConfig {
			Host: getEnv("DB_HOST", "localhost"),
			Port: dbPort,
			User: getEnv("DB_USER", "postgres"),
			Pass: getEnv("DB_PASSWORD", "pass"),
			Name: getEnv("DB_NAME", "mydb"),
		},
	}
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}
