package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	Item      string `gorm:"type:varchar(255)" json:"item"`
	Completed bool   `gorm:"type:varchar(255)" json:"completed"`
}
