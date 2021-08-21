package response


var MsgFlags =map[int]string{
	ERROR_AUTH_PARAM_FAIL: "登录参数错误，请检查登陆参数",
	ERROR_DATABASE_QUERY:"数据库查询错误，请联系管理员",
	ERROR_TOKEN_GENERATE_FAIL:"token生成错误",
	ERROR_ADMIN_INVALID_PASSWORD:"管理员密码错误",
	ERROR_LOGIN_USERNAME:"用户名格式错误",
	ERROR_AUTH_UNKNOWN_ERROR:"用户名未找到",
	ERROR_AUTH_INVALID_PASSWORD:"密码错误",
	ERROR_AUTH_NO_VALID_HEADER:"请求头格式错误，请检查Authorization字段",
	ERROR_DEFAULT: "未知错误",

}

func GetMsg(code int) string {
	msg,ok:=MsgFlags[code]
	if ok{
		return msg
	}
	return MsgFlags[ERROR_DEFAULT]
}