package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppEnv struct {
	Host         string
	Port         string
	Password     string
	User         string
	SSLMode      string
	DbName       string
	DbUrl        string
	JwtSecretKey []byte
}

func SetupEnv() *AppEnv {
	if err := godotenv.Load(".env"); err != nil {
		println("Could not load env")
	}

	jwt := os.Getenv("JWT_SECRET_KEY")
	config := &AppEnv{
		DbUrl:        os.Getenv("DB_URL"),
		JwtSecretKey: []byte(jwt),
	}
	return config
}
