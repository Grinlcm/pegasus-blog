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

//type Blog struct {
//	ID          uint     `gorm:"primary_key"`
//	Title       string   `gorm:"not null;size:256"`
//	Content     string   `gorm:"type:text;not null"`
//	ShowContent string   `gorm:"not null"`
//	LookNum     int      `gorm:"default:1"`
//	BlogType    BlogType `gorm:"ForeignKey:BlogTypeID;AssociationForeignKey:ID"`
//	//这个AssiciationForignKey 可以指明BlogType中的哪个字段当做 关联的字段， 可以将ID改成Name
//	// 这样关联的字段就是Name了
//	BlogTypeID int // 还需要写上这
//}
//
//type BlogType struct {
//	ID   uint `gorm:"primary_key"`
//	Name string
//}
