package utils

func IsAllowPwd(pwd string) bool {
	return IsAllowLen(pwd, 6, 50)
}

func IsAllowLen(s string, minLen, maxLen int) bool {
	n := len(s)
	return n >= minLen && n <= maxLen
}
