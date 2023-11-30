package global

import (
	"github.com/Mmx233/EnvConfig"
	"strings"
)

type _Config struct {
	Addr string
	// default 30s
	Timeout uint `config:"omitempty"`

	// default 7d
	LoginValidate uint   `config:"omitempty"`
	WhiteListPath string `config:"omitempty"`

	AppCode   string
	AppSecret string
}

var Config _Config

var WhiteListPath []string

func initConfig() {
	EnvConfig.Load("", &Config)

	if Config.Timeout == 0 {
		Config.Timeout = 30
	}
	if Config.LoginValidate == 0 {
		Config.LoginValidate = 7
	}

	if Config.WhiteListPath != "" {
		WhiteListPath = strings.Split(strings.TrimSpace(Config.WhiteListPath), ",")
	}
}
