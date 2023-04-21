package global

import (
	"github.com/Mmx233/EnvConfig"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global/models"
	"strings"
)

var Config models.Config

var AllowGroups []string

func initConfig() {
	EnvConfig.Load("", &Config)

	if Config.Timeout == 0 {
		Config.Timeout = 30
	}
	if Config.Groups != "" {
		AllowGroups = strings.Split(strings.TrimSpace(Config.Groups), ",")
	}
}
