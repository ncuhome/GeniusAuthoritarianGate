package global

import (
	"strings"
)

type _Config struct {
	GeniusAuthHost string `config:"omitempty"`

	Addr string
	// default 30s
	Timeout uint `config:"omitempty"`
	AesKey  string

	// default 7d
	LoginValidate uint   `config:"omitempty"`
	WhiteListPath string `config:"omitempty"`

	AppCode   string
	AppSecret string
}

var Config _Config

var WhiteListPath []string

func fillDefaultConfig() {
	if Config.GeniusAuthHost == "" {
		Config.GeniusAuthHost = "v.ncuos.com"
	}

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
