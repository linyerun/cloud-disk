package utils

import (
	"cloud-disk/core/define"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"regexp"
	"strconv"
	"time"
)

// IsNormalEmail 验证邮箱
func IsNormalEmail(email string) bool {
	// 最大长度不能超过320
	if len(email) > 320 {
		return false
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

// GetCode 生成验证码
func GetCode() string {
	return RandomCode(6)
}

// RandomCode 生成长度为n的随机数
func RandomCode(n int) string {
	rand.Seed(time.Now().Unix()) // 种随机种子
	res := ""
	for i := 0; i < n; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}

// SendCode 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "CloudDesk <2338244917@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	return e.Send("smtp.qq.com:25", smtp.PlainAuth("", "2338244917@qq.com", define.MailPassword, "smtp.qq.com"))
}
