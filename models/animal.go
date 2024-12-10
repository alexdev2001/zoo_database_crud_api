package models

type Animal struct {
	Id       uint   `gorm:"primary_key"`
	Type     string `gorm:"size:20"`
	Quantity uint   `gorm:"size:20"`
}
