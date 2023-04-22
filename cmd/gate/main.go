package main

import (
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/api/gateway"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infoln("Sys Boost")
	if e := gateway.Run(":80"); e != nil {
		log.Fatalln("启动监听失败:", e)
	}
}
