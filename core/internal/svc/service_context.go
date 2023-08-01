package svc

import (
	"cloud-disk/core/db"
	"cloud-disk/core/internal/config"
	"cloud-disk/core/internal/middleware"
	"cloud-disk/core/models"
	"cloud-disk/core/process"
	"cloud-disk/core/utils"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config                 config.Config
	MySQLClient            *gorm.DB
	RedisClient            *redis.Client
	Auth                   rest.Middleware
	SaveOrRejectRemoteAddr rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	ret := &ServiceContext{
		Config:                 c,
		MySQLClient:            db.InitMySQL(&c),
		RedisClient:            db.InitRedis(&c),
		Auth:                   middleware.NewAuthMiddleware().Handle,
		SaveOrRejectRemoteAddr: middleware.NewSaveOrRejectRemoteAddrMiddleware().Handle,
	}
	ret.generateTable() // 如果没有创建表使用这个函数来生成
	process.Init()      // 初始化协程池
	return ret
}

func (ret *ServiceContext) generateTable() {
	if err := ret.MySQLClient.AutoMigrate(models.User{}, models.File{}, models.UserFile{}, models.ShareFile{}); err != nil {
		utils.Logger().Error(err)
	}
}
