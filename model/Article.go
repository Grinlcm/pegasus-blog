package model

import (
	"gorm.io/gorm"
	"pegasus-blog/util/errmsg"
)

type Article struct {
	ID         uint     `gorm:"primary_key" json:"ID"`
	Title      string   `gorm:"type:varchar(100);not null" json:"title"`
	Desc       string   `gorm:"type:varchar(200);" json:"desc"`
	Content    string   `gorm:"type:longtext" json:"content"`
	Img        string   `gorm:"type:varchar(100)" json:"img"`
	Category   Category `gorm:"ForeignKey:CategoryID;AssociationForeignKey:ID"`
	CategoryID int      `gorm:"type:int;not null" json:"category_id"`
}

// CreateArticle  新增文章
func CreateArticle(data *Article) int {
	err = DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArticle 查询文章
func GetArticle(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err = DB.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCESS
}

// 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err = DB.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERRORArticleNotExist
	}
	return article, errmsg.SUCCESS
}

// 查询分类下所有文章
func GetCateArticle(id int, pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err := DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("category_id = ?", id).Find(&articleList).Error
	if err != nil {
		return nil, errmsg.ERRORCategoryNotExist
	}
	return articleList, errmsg.SUCCESS
}

// EditArticle 编辑文章
func EditArticle(id int, data *Article) int {
	var updataMap = make(map[string]interface{})
	updataMap["title"] = data.Title
	updataMap["desc"] = data.Desc
	updataMap["content"] = data.Content
	updataMap["img"] = data.Img
	updataMap["category_id"] = data.CategoryID

	err := DB.Model(&Article{}).Where("id = ?", id).Updates(updataMap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var art Article
	err = DB.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
