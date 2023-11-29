package global

import (
	"github.com/Mmx233/EnvConfig"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global/models"
	"strings"
)

var Config models.Config

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
