package global

import (
	"github.com/Mmx233/EnvConfig"
	"github.com/Mmx233/tool"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global/models"
	"strings"
)

var Config models.Config

var AllowGroups []string

var WhiteListPath []string

func initConfig() {
	EnvConfig.Load("", &Config)

	if Config.Timeout == 0 {
		Config.Timeout = 30
	}
	if Config.LoginValidate == 0 {
		Config.LoginValidate = 7
	}
	if Config.JwtKey == "" {
		Config.JwtKey = tool.Rand.String(64)
	}

	if Config.Groups != "" {
		AllowGroups = strings.Split(strings.TrimSpace(Config.Groups), ",")
	}
	if Config.WhiteListPath != "" {
		WhiteListPath = strings.Split(strings.TrimSpace(Config.WhiteListPath), ",")
	}
}
