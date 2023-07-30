package define

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
)

const (
	RedisCodePrefix = "code_"
	JwtKey          = "cloud-desk-lyr-goZero-jwt"
	MaxFileSize     = 500 * 1024 * 1024 // 单位: B
)

const (
	EachCodeForEmailWaitTime = 20 * 60 // 单位: 秒, 验证码发送一次用户需要等待的时间
	TokenExpireTime          = 30 * 60 // 单位: 秒，token过期时间
)

var (
	MailPassword     = os.Getenv("MailPassword")
	CosBucket        = os.Getenv("CosBucket")
	TencentSecretID  = os.Getenv("TencentSecretID")
	TencentSecretKey = os.Getenv("TencentSecretKey")
)

type UserClaim struct {
	ID    uint
	Email string
	jwt.RegisteredClaims
}
