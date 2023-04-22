package gateway

import (
	gateway "github.com/Mmx233/Gateway/v2"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/api/middlewares"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Run(addr ...string) error {
	gin.SetMode(gin.ReleaseMode)
	E := gin.Default()

	E.Use(middlewares.Auth(global.Config.JwtKey))

	E.Use(gateway.Proxy(&gateway.ApiConf{
		Addr:      global.Config.Addr,
		Transport: util.HttpTransport,
		ErrorHandler: func(_ http.ResponseWriter, _ *http.Request, e error) {
			log.Errorln("request backend failed:", e)
		},
		AllowAll: true,
	}))

	return E.Run(addr...)
}
