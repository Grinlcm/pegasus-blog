package model

type Article struct {
	ID         uint     `gorm:"primary_key" json:"ID"`
	Title      string   `gorm:"type:varchar(100);not null" json:"title"`
	Desc       string   `gorm:"type:varchar(200);" json:"desc"`
	Content    string   `gorm:"type:longtext" json:"content"`
	Img        string   `gorm:"type:varchar(100)" json:"img"`
	Category   Category `gorm:"ForeignKey:CategoryID;AssociationForeignKey:ID"`
	CategoryID int
}
