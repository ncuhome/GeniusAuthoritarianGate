package ga

import (
	geniusAuth "github.com/ncuhome/GeniusAuthoritarianClient"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	log "github.com/sirupsen/logrus"
)

var Client *geniusAuth.Client
var Rpc *geniusAuth.RpcClient

func init() {
	Client = geniusAuth.NewClient(
		global.Config.GeniusAuthHost,
		global.Config.AppCode, global.Config.AppSecret,
		util.Http.Client,
	)

	var err error
	Rpc, err = Client.NewRpcClient(global.Config.GeniusAuthAppRpcAddr)
	if err != nil {
		log.Fatalln("Create GeniusAuth rpc connection failed:", err)
	}
}
