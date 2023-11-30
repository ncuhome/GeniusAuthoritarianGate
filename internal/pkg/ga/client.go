package ga

import (
	geniusAuth "github.com/ncuhome/GeniusAuthoritarianClient"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
)

var Client *geniusAuth.Client

func init() {
	Client = geniusAuth.NewClient(
		global.Config.GeniusAuthHost,
		global.Config.AppCode, global.Config.AppSecret,
		util.Http.Client,
	)
}
