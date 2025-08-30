package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	Port       string
}

func Load() *Config {
	env := getEnv("APP_ENV", "prod")

	projectRoot, _ := os.Getwd()
	envFile := fmt.Sprintf("%s/profiles/%s.env", projectRoot, env)
	log.Printf("Load envFile: %s", envFile)

	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("failed to load %s file: %v", envFile, err)
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "user"),
		DBUser:     getEnv("DB_USER", "jey_user"),
		DBPassword: getEnv("DB_PASS", "pass_p@22w@rd"),
		Port:       getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
