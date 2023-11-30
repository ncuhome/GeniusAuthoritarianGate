package middlewares

import (
	webServe "github.com/Mmx233/GinWebServe"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/GeniusAuthoritarianGate/web"
	log "github.com/sirupsen/logrus"
	"strings"
)

func LoginWeb() gin.HandlerFunc {
	fs, err := web.Fs()
	if err != nil {
		log.Fatalln(err)
	}

	handler, err := webServe.New(fs)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/login/") && !strings.HasPrefix(c.Request.URL.Path, "/login/api/") {
			handler(c)
		}
	}
}
