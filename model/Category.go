package model

import (
	"pegasus-blog/util/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primary_key;autoincrement" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) int {
	var cate Category
	DB.Select("id").Where("name = ? ", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ErrorCategoryUsed
	}
	return errmsg.SUCCESS
}

// CreateCategory  新增分类
func CreateCategory(data *Category) int {
	err = DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCategory 查询分类
func GetCategory(pageSize int, pageNum int) []Category {
	var cates []Category
	err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil {
		return nil
	}
	return cates
}

// EditCategory 编辑分类
func EditCategory(id int, data *Category) {
	var updataMap = make(map[string]interface{})
	updataMap["name"] = data.Name
	DB.Model(&Category{}).Where("id = ?", id).Updates(updataMap)

}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err = DB.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
