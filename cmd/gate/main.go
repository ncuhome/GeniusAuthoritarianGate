package main

import (
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/api/gateway"
	"github.com/sirupsen/logrus"
)

func main() {
	if e := gateway.Run(":80"); e != nil {
		logrus.Fatalln("启动监听失败:", e)
	}
}
