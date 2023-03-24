package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// ErrorUserUsed code=1000 用户模块的错误
	ErrorUserUsed       = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenExist     = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	// ErrorCategoryUsed code=2000 文章模块的错误

	ERRORArticleNotExist = 2001
	// ErrorCategoryUsed code=3000 分类模块的错误
	ErrorCategoryUsed     = 3001
	ERRORCategoryNotExist = 3002
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "失败",

	ErrorUserUsed:       "用户名已存在",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExist:   "用户不存在",
	ErrorTokenExist:     "TOKEN不存在",
	ErrorTokenRuntime:   "TOKEN已过期",
	ErrorTokenWrong:     "TOKEN不正确",
	ErrorTokenTypeWrong: "TOKEN格式不正确",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
