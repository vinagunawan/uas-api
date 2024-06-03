package entities

import "time"

type WaterIntake struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Timestamp time.Time
}
