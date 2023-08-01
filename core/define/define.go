package define

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
)

const (
	RedisCodePrefix = "code_"
	JwtKey          = "cloud-desk-lyr-goZero-jwt"
	MaxFileSize     = 500 * 1024 * 1024 // 单位: B
	Host            = "https://cloud.disk.daluotang.cc"
	ShareFilePrefix = "share_file:"
	UpdateClickNum  = 200 // 每间隔多少次更新点赞才保存到数据库
)

const (
	EachCodeForEmailWaitTime = 20 * 60 // 单位: 秒, 验证码发送一次用户需要等待的时间
	TokenExpireTime          = 30 * 60 // 单位: 秒，token过期时间
	CacheExpireTime          = 10 * 60 // 单位: 秒，redis缓存过期时间
	ProcessingCenterPoolSize = 10      // 协程池大小
	ProcessingCenterChanLen  = 10      //协程任务队列大小
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
