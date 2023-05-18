package verify

import "regexp"

// 邮箱脱敏
func HideEmail(email string) (result string) {
	pattern := `(\w?)(\w+)(\w)(@\w+\.[a-z]+(\.[a-z]+)?)`
	reg := regexp.MustCompile(pattern)
	return reg.ReplaceAllString(email, "$1****$3$4")
}

// 手机号脱敏
func HidePhoneNumber(phone string) string {
	pattern := `(\d{3})(\d*)(\d{4})`
	reg := regexp.MustCompile(pattern)
	return reg.ReplaceAllString(phone, "$1****$3")
}

// 邮箱验证
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	//pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 手机号验证
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
