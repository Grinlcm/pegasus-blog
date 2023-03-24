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
		// 分类模块的接口
		router.POST("/category/add", v1.AddCategory)
		router.GET("/category", v1.GetCategory)
		router.PUT("/category/:id", v1.EditCategory)
		router.DELETE("/category/:id", v1.DeleteCategory)
		// 文章模块的接口
		router.POST("/article/add", v1.AddArticle)
		router.GET("/article", v1.GetArticle)              // 查询文章列表
		router.GET("/article/info/:id", v1.GetArticleInfo) // 查询单个文章
		router.GET("/article/list", v1.GetCategoryArticle) // 查询单个分类下的所有文章

		router.PUT("/article/:id", v1.EditArticle)
		router.DELETE("/article/:id", v1.DeleteArticle)
		// 登录模块的接口
	}

	r.Run(fmt.Sprintf(":%s", util.Appconfig.HttpPort))
}
