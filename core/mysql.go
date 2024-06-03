package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := GetEnv("DB_USER") + ":" + GetEnv("DB_PASS") + "@tcp(" + GetEnv("DB_HOST") + ":" + GetEnv("DB_PORT") + ")/" + GetEnv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
