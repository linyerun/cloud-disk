package db

import (
	"cloud-disk/core/internal/config"
	"cloud-disk/core/utils"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path"
	"time"
)

var RedisClient *redis.Client
var MySQLClient *gorm.DB

const RedisNil = "redis: nil"

func InitRedis(cnf *config.Config) *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cnf.Redis.Host, cnf.Redis.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := db.Ping(context.Background()).Err(); err != nil {
		utils.Logger().Warn("请开启Redis服务端，默认地址为: localhost:6379")
		utils.Logger().Error(err)
		panic(err)
	}
	RedisClient = db
	return db
}

func InitMySQL(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local&timeout=10s&readTimeout=10s&writeTimeout=10s", cnf.MySQL.Username, cnf.MySQL.Password, cnf.MySQL.Host, cnf.MySQL.Port, cnf.MySQL.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true, //预编译语句
		Logger:                 newGormLogger(),
	})
	if err != nil {
		utils.Logger().Error(err)
		panic(err)
	}
	MySQLClient = db
	return db
}

func InitMySQLForTest(Username, Password, Host, Database string, Port uint) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local&timeout=10s&readTimeout=10s&writeTimeout=10s", Username, Password, Host, Port, Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true, //预编译语句
		Logger:                 newGormLogger(),
	})
	if err != nil {
		utils.Logger().Error(err)
		panic(err)
	}
	MySQLClient = db
}

func newGormLogger() logger.Interface {
	//创建文件夹
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileDir := rootDir + "/cloud_disk_logs"
	if _, err = os.Stat(fileDir); os.IsNotExist(err) {
		//不存在这个文件夹就创建
		err := os.MkdirAll(fileDir, 0666) //可读写，不可执行
		if err != nil {
			panic(err)
		}
	}

	//创建文件
	fileName := "gorm.log"
	filePath := path.Join(fileDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	return logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,
		},
	)
}
