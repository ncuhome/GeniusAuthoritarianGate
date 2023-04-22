package gateway

import (
	gateway "github.com/Mmx233/Gateway"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/api/middlewares"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	log "github.com/sirupsen/logrus"
)

func Run(addr ...string) error {
	gin.SetMode(gin.ReleaseMode)
	E := gin.Default()

	E.Use(middlewares.Auth(global.Config.JwtKey))

	E.Use(gateway.Proxy(&gateway.ApiConf{
		Addr:   global.Config.Addr,
		Client: util.Http.Client,
		ErrorHandler: func(c *gin.Context, e error) {
			log.Errorln("request backend failed:", e)
			c.AbortWithStatus(502)
		},
		AllowAll: true,
	}))

	return E.Run(addr...)
}
