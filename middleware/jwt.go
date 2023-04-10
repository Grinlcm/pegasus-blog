package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"pegasus-blog/util"
	"pegasus-blog/util/errmsg"
	"strings"
	"time"
)

var JwtKey = []byte(util.Appconfig.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// SetToken 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(time.Hour * 10).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "pegasus-blog",
		},
	})
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return tokenString, errmsg.SUCCESS
}

// VerifyToken 验证token
func VerifyToken(tokenString string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

// jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		if tokenHeader == "" {
			code = errmsg.ErrorTokenExist
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if (len(checkToken) != 2) || (checkToken[0] != "Bearer") {
			code = errmsg.ErrorTokenTypeWrong
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, checkCode := VerifyToken(checkToken[1])
		if checkCode == errmsg.ERROR {
			code = errmsg.ErrorTokenWrong
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ErrorTokenRuntime
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}
}
