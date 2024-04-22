package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	DBAddress  string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "yasba_user"),
		DBPassword: getEnv("DB_PASSWORD", "yasba_password"),
		DBName:     getEnv("DB_NAME", "yasba_db"),
		JWTSecret:  getEnv("JWT_SECRET", "randomjwtsecretkey"),
		DBAddress:  fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", getEnv("DB_USER", "yasba_user"), getEnv("DB_PASSWORD", "yasba_password"), getEnv("DB_NAME", "yasba_db")),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
