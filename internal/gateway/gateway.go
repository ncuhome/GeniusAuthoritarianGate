package gateway

import (
	gateway "github.com/Mmx233/Gateway/v2"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/api/controllers"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/middlewares"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Run(addr ...string) error {
	gin.SetMode(gin.ReleaseMode)
	E := gin.Default()

	E.Use(middlewares.Auth(), middlewares.LoginWeb())

	api := E.Group("login/api")
	api.POST("/", controllers.Login)

	E.Use(gateway.Proxy(&gateway.ApiConf{
		Addr:      global.Config.Addr,
		Transport: util.Http.Client.Transport,
		ErrorHandler: func(_ http.ResponseWriter, _ *http.Request, e error) {
			log.Errorln("request backend failed:", e)
		},
		AllowAll: true,
	}))

	return E.Run(addr...)
}
