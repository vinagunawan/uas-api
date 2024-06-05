package entities

import "time"

const TableNameWaterIntake = "water_intakes"

type WaterIntake struct {
	ID     int       `gorm:"column:id;primaryKey" json:"id"`
	UserID int       `gorm:"column:user_id" json:"user_id"`
	Date   time.Time `gorm:"column:date" json:"date"`
	Amount float64   `gorm:"column:amount" json:"amount"`
}

func (*WaterIntake) TableName() string {
	return TableNameWaterIntake
}
