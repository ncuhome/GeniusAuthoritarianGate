package main

import (
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/gateway"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infoln("Sys Boost")
	if err := gateway.Run(":80"); err != nil {
		log.Fatalln("listen ::80 failed:", err)
	}
}
