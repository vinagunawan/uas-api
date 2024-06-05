package core

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(app *Application) *gorm.DB {
	db, err := gorm.Open(mysql.Open(app.Env.MysqlUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	return db
}
