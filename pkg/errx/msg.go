package errx

var message = map[uint32]string{
	Ok:             "SUCCESS",
	ERROR:          "Fail",
	NOT_AUTH:       "没有认证",
	PARAM_ERROR:    "请求参数错误",
	INVALID_PARAMS: "请求参数格式错误",
	NOT_FOUND:      "找不到指定的资源",
	SERVER_ERROR:   "服务器开小差啦",
	DB_ERROR:       "数据库繁忙",
	UNKNOWN_ERROR:  "未知错误",

	USERNAME_ALREADY_EXISTS:       "用户名已存在",
	USER_LOGIN_ERROR:              "用户名或密码错误",
	PASSWORLD_ERROR:               "密码错误",
	INVALID_PASSWORD_FORMAT_ERROR: "密码格式错误",
	EMAIL_ALREADY_REGISTER_ERROR:  "邮箱已注册",
	EMAIL_NOT_REGISTER_ERROR:      "邮箱未注册",
	USER_NOT_FOUND_ERROR:          "用户不存在",

	INVALID_USERNAME_FORMAT_ERROR: "用户名格式错误",
	INVALID_EMAIL_FORMAT_ERROR:    "邮箱格式错误",
}

func GetMessage(code uint32) string {
	msg, ok := message[code]
	if ok {
		return msg
	}

	return message[UNKNOWN_ERROR]
}
