package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Mysql  `toml:"mysql"`
	Base   `toml:"base"`
	WeChat `toml:"wechat"`
}

type Base struct {
	Port string
}

type Mysql struct {
	DBHost     string
	DBUsername string
	DBPassword string
	DBName     string
}

type WeChat struct {
	CorpID         string
	CorpSecret     string
	AccessTokenUrl string
	AdminOpenID    string
	SendMsgUrl     string
	AgentID        int
}

func (s *Service) initConfig() {
	var conf *Config
	_, err := toml.DecodeFile("./conf/config.toml", &conf)
	if err != nil {
		log.Fatalln(err)
	}

	s.Conf = conf
	log.Println("config load success")
}
