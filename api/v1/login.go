package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pegasus-blog/middleware"
	"pegasus-blog/model"
	"pegasus-blog/util/errmsg"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)

	var token string
	var code int
	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
