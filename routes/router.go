package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pegasus-blog/util"
)

func InitRouter() {
	gin.SetMode(util.Appconfig.AppMode)

	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "你好",
			})
		})
	}

	r.Run(fmt.Sprintf(":%s", util.Appconfig.HttpPort))
}
