package core

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port     string
	MysqlUrl string
}

func NewEnv() *Env {
	godotenv.Load()

	return &Env{
		Port:     os.Getenv("PORT"),
		MysqlUrl: os.Getenv("MYSQL_URL"),
	}
}
