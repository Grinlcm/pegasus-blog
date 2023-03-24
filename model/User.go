package model

import (
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"pegasus-blog/util/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"  json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

func CheckUserName(name string) int {
	var users User
	DB.Select("id").Where("username = ? ", name).First(&users)
	if users.ID > 0 {
		return errmsg.ErrorUserUsed
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	err = DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 查询用户
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password, err = ScryptPw(u.Password)
	if err != nil {
		return errors.New("加密错误")
	}
	return nil
}

// EditUser 编辑用户
func EditUser(id int, data *User) {
	var updataMap = make(map[string]interface{})
	updataMap["name"] = data.Username
	updataMap["role"] = data.Role
	DB.Model(&User{}).Where("id = ?", id).Updates(updataMap)

}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = DB.Where("id = ?").Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func ScryptPw(password string) (string, error) {
	var KeyLen = 32
	salt := make([]byte, 8)
	salt = []byte{12, 15, 17, 89, 134, 234, 246, 78}
	hashPWD, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, KeyLen)
	if err != nil {
		fmt.Println("加密报错")
	}
	return base64.StdEncoding.EncodeToString(hashPWD), err
}
