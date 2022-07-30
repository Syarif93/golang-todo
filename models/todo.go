package models

type Todo struct {
	// gorm.Model
	ID        string `gorm:"primary_key:auto_increment" json:"id"`
	Item      string `gorm:"type:varchar(255)" json:"item"`
	Completed bool   `gorm:"type:varchar(255)" json:"completed"`
}
