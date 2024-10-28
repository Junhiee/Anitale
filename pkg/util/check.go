package util

import "regexp"

// CheckEmail 校验邮箱格式
func CheckEmail(email string) bool {
	// 定义一个正则表达式用于邮箱格式校验
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// CheckUsername 校验用户名
func CheckUsername(username string) bool {
	// 用户名长度限制
	if len(username) < 3 || len(username) > 20 {
		return false
	}

	// 检查用户名是否只包含字母、数字、下划线
	const usernameRegex = `^[a-zA-Z0-9_]+$`
	re := regexp.MustCompile(usernameRegex)
	return re.MatchString(username)
}

// CheckPassword 校验密码复杂性
// 密码长度必须大于等于6
// 密码必须包含一个字母和一个数字
func CheckPassword(pwd string) bool {
	// 检查密码长度
	if len(pwd) < 6 {
		return false
	}

	// 定义正则表达式
	var (
		hasLetter = regexp.MustCompile(`[a-zA-Z]`)
		hasNumber = regexp.MustCompile(`[0-9]`)
		// hasSpecial   = regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	)

	// 检查密码是否包含一个字母和一个数字
	return hasLetter.MatchString(pwd) &&
		hasNumber.MatchString(pwd)
}
