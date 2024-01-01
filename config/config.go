package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct{
	Username string
	Password string
	DBName string
	Host string
	Port string

}

func Load(){
	if err := godotenv.Load(); err != nil{
		log.Println("No .env file found!")
	}
}

func GetDatabaseConfig() DatabaseConfig{
	return DatabaseConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
	}
}

func GetSecretKey() string{
	return os.Getenv("SECRET_KEY")
}