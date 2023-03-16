package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "pegasus-blog/api/v1"
	"pegasus-blog/util"
)

func InitRouter() {
	gin.SetMode(util.Appconfig.AppMode)

	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 用户模块的接口
		router.POST("/user/add", v1.AddUser)
		router.GET("/users", v1.GetUsers)
		router.PUT("/user/:id", v1.EditUser)
		router.DELETE("/user/:id", v1.DeleteUser)
		// 文章模块的接口
		// 登录模块的接口
	}

	r.Run(fmt.Sprintf(":%s", util.Appconfig.HttpPort))
}
