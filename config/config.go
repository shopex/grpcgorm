package config

import (
	"log"

	"git.shopex.cn/luban/tools/conf"
)

var Config *config

func init() {
	Config = &config{}
	err := conf.Load("env.conf", Config)
	if err != nil {
		log.Println(err)
		conf.Load("../env.conf", Config)
	}
}

type config struct {
	DataBase string `cfg:"DATABASE"`
	Port     string `cfg:"PORT"`
}
