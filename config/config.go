package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	DbHostPsql     string
	DbPortPsql     string
	DbUserNamePsql string
	DbPasswordPsql string
	DbPsql         string
	AuthSecretKey  string
}

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("error while load .env file")
	}
}

func GetConfig() AppConfig {
	return AppConfig{
		DbHostPsql:     os.Getenv("HOST_PSQL"),
		DbPortPsql:     os.Getenv("PORT_PSQL"),
		DbUserNamePsql: os.Getenv("USERNAME_PSQL"),
		DbPasswordPsql: os.Getenv("PASSWORD_PSQL"),
		DbPsql:         os.Getenv("DB_PSQL"),
		AuthSecretKey:  os.Getenv("AUTH_SECRET_KEY"),
	}
}
