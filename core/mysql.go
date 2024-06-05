package core

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(app *Application) *gorm.DB {
	log.Printf("Connecting to database with URL: %s", app.Env.MysqlUrl)
	db, err := gorm.Open(mysql.Open(app.Env.MysqlUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Successfully connected to database")

	return db
}
