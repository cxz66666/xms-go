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
	ERROR_NOT_LOGIN:"未登录，请先登录",
	ERROR_TOKEN_NOT_VAILD:"令牌不合法，请重新登录",
	ERROR_TOKEN_EXPIRED:"token令牌已过期,请重新获取",
	ERROR_IP_BLOCK:"登录失败次数过多，ip暂时封禁，请5分钟后再尝试",
	ERROR_FORBIDDEN:"您的账号权限不足，请联系管理员授权",
	ERROR_BILL_INVALID_TYPE:"账单表单参数错误",
	ERROR_PARAM_NOT_VAILD:"param参数不合法",
	ERROR_BILL_NOT_FOUND:"账单未找到",
	ERROR_BILL_INVALID_QUERY:"query参数错误",
	ERROR_NEWS_NOT_FOUND:"news没有找到",
	ERROR_NEWS_INVALID_INFO:"news模型绑定错误，参数不合法",
	ERROR_NEWS_INVALID_QUERY:"query参数错误，请检查参数是否合法",
	ERROR_TICKET_SEARCH_INVALID_QUERY:"query参数错误，请检查参数是否合法",
	ERROR_TICKET_LIST_INVALID_QUERY:"query参数错误，请检查参数是否合法",
	ERROR_TICKET_NOT_FOUND:"ticket未找到，请检查id",
	ERROR_TICKET_INVALID_INFO:"ticket模型绑定失败",
	ERROR_TICKET_IS_CONFIRMED:"ticket已经锁定",
	ERROR_TICKET_ALREADY_PICKED:"ticket已经认领过了",
	ERROR_NOT_LOCKABLE:"ticket无法被锁定",
	ERROR_TICKET_NOT_CONFIRMED:"ticket未被锁定",
	ERROR_USER_NOT_FOUND:"用户未找到",
	ERROR_DEFAULT: "未知错误",

}

func GetMsg(code int) string {
	msg,ok:=MsgFlags[code]
	if ok{
		return msg
	}
	return MsgFlags[ERROR_DEFAULT]
}