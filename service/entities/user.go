package entities

const TableNameUser = "users"

type User struct {
	ID       int     `gorm:"column:id;primaryKey" json:"id"`
	Username string  `gorm:"column:username" json:"username"`
	Password string  `gorm:"column:password" json:"password"`
	Weight   float64 `gorm:"column:weight" json:"weight"`
	Height   float64 `gorm:"column:height" json:"height"`
	Age      int     `gorm:"column:age" json:"age"`
}

func (*User) TableName() string {
	return TableNameUser
}
