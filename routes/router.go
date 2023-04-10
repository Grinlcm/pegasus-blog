package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "pegasus-blog/api/v1"
	"pegasus-blog/middleware"
	"pegasus-blog/util"
)

func InitRouter() {
	gin.SetMode(util.Appconfig.AppMode)

	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtMiddleware())
	{
		// 用户模块的接口

		auth.GET("/users", v1.GetUsers)
		auth.PUT("/user/:id", v1.EditUser)
		auth.DELETE("/user/:id", v1.DeleteUser)
		// 分类模块的接口
		auth.POST("/category/add", v1.AddCategory)

		auth.PUT("/category/:id", v1.EditCategory)
		auth.DELETE("/category/:id", v1.DeleteCategory)
		// 文章模块的接口
		auth.POST("/article/add", v1.AddArticle)

		auth.PUT("/article/:id", v1.EditArticle)
		auth.DELETE("/article/:id", v1.DeleteArticle)
		// 登录模块的接口
	}

	public := r.Group("/api/v1")
	{
		public.POST("/user/add", v1.AddUser)

		public.GET("/category", v1.GetCategory)
		public.GET("/article", v1.GetArticle)              // 查询文章列表
		public.GET("/article/info/:id", v1.GetArticleInfo) // 查询单个文章
		public.GET("/article/list", v1.GetCategoryArticle) // 查询单个分类下的所有文章
		public.POST("/login", v1.Login)
	}

	r.Run(fmt.Sprintf(":%s", util.Appconfig.HttpPort))
}
