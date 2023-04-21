package gateway

import (
	gateway "github.com/Mmx233/Gateway"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	"github.com/sirupsen/logrus"
)

func Run(addr ...string) error {
	gin.SetMode(gin.ReleaseMode)
	E := gin.Default()

	E.Use(gateway.Proxy(&gateway.ApiConf{
		Addr:   global.Config.Addr,
		Client: util.Http.Client,
		ErrorHandler: func(c *gin.Context, e error) {
			logrus.Errorln("request backend failed:", e)
			c.AbortWithStatus(502)
		},
		AllowAll: true,
	}))

	return E.Run(addr...)
}
