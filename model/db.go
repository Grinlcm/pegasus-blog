package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pegasus-blog/util"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		util.DBConfig.UserName,
		util.DBConfig.Password,
		util.DBConfig.Host,
		util.DBConfig.Port,
		util.DBConfig.DBName,
		util.DBConfig.Charset,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("连接数据库失败, 请检查参数：%v", err)
	}

	DB.AutoMigrate(&User{}, Article{}, Category{})

	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Printf("连接池设置失败:%v", err)
	}

	// SetMaxIdleConns 最大闲置连接数
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 最大连接数
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
