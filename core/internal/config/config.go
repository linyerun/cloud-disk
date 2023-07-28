package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MySQL struct {
		Username string
		Password string
		Host     string
		Database string
		Port     uint
	}
	Redis struct {
		Host string
		Port uint
	}
}
