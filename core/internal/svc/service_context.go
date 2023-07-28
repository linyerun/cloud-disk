package svc

import (
	"cloud-disk/core/db"
	"cloud-disk/core/internal/config"
	"cloud-disk/core/models"
	"cloud-disk/core/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	MySQLClient *gorm.DB
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	ret := &ServiceContext{
		Config:      c,
		MySQLClient: db.InitMySQL(&c),
		RedisClient: db.InitRedis(&c),
	}
	ret.generateTable() // 如果没有创建表使用这个函数来生成
	return ret
}

func (ret *ServiceContext) generateTable() {
	if err := ret.MySQLClient.AutoMigrate(models.User{}); err != nil {
		utils.Logger().Error(err)
	}
}
