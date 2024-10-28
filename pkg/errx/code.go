package errx

const (
	// 全局错误码
	Ok    uint32 = 200
	ERROR uint32 = 502

	NOT_AUTH       uint32 = 401
	PARAM_ERROR    uint32 = 400
	INVALID_PARAMS uint32 = 422 // 参数格式错误
	NOT_FOUND      uint32 = 404

	SERVER_ERROR  uint32 = 500
	DB_ERROR      uint32 = 501
	UNKNOWN_ERROR uint32 = 999

	// 业务错误码
	USER_NOT_FOUND_ERROR          uint32 = 300008 // 用户不存在
	USERNAME_ALREADY_EXISTS       uint32 = 300003 // 用户名已存在
	INVALID_USERNAME_FORMAT_ERROR uint32 = 300002 // 用户名格式错误
	USER_LOGIN_ERROR              uint32 = 300005 // 用户名或密码错误

	PASSWORLD_ERROR               uint32 = 300006 // 密码错误
	INVALID_PASSWORD_FORMAT_ERROR uint32 = 300009 // 密码格式错误
	EMAIL_ALREADY_REGISTER_ERROR  uint32 = 300004 // 邮箱已注册
	EMAIL_NOT_REGISTER_ERROR      uint32 = 300007 // 邮箱未注册
	INVALID_EMAIL_FORMAT_ERROR    uint32 = 300001 // 邮箱格式错误

)
