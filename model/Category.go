package model

type Category struct {
	ID   uint   `gorm:"primary_key" json:"ID"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
