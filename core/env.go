package core

import (
	"log"
	"os"
	
	"github.com/joho/godotenv"
)

type Env struct {
	Port     string
	MysqlUrl string
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	mysqlUrl := os.Getenv("MYSQL_URL")

	if port == "" || mysqlUrl == "" {
		log.Fatalf("Required environment variables PORT or MYSQL_URL are missing")
	}

	log.Printf("Loaded environment variables: PORT=%s, MYSQL_URL=%s", port, mysqlUrl)

	return &Env{
		Port:     port,
		MysqlUrl: mysqlUrl,
	}
}
