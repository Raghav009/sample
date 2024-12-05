package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	Database      string
}

func LoadConfig() (*Config, error) {
	// curDir, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }
	path := "D:\\Training\\sample\\.env" // curDir + "/.env" //
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		Database:      getEnv("DATABASE_URL", ""),
	}, nil
}

func LoadSecret() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	return getEnv("SECRET_KEY", ""), nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
