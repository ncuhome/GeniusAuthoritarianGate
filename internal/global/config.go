package global

import (
	"github.com/Mmx233/EnvConfig"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global/models"
)

var Config models.Config

func initConfig() {
	EnvConfig.Load("", &Config)
}
