package entities

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Weight   float64
	Height   float64
	Age      int
}
