package errx

var message = map[uint32]string{
	Ok:            "SUCCESS",
	ERROR:         "Fail",
	NOT_AUTH:      "没有认证",
	PARAM_ERROR:   "请求参数错误",
	NOT_FOUND:     "找不到指定的资源",
	SERVER_ERROR:  "服务器开小差啦",
	DB_ERROR:      "数据库繁忙",
	UNKNOWN_ERROR: "未知错误",
}

func GetMessage(code uint32) string {
	msg, ok := message[code]
	if ok {
		return msg
	}

	return message[UNKNOWN_ERROR]
}
