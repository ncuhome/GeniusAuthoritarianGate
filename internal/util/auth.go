package util

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
)

func GoGeniusLogin(c *gin.Context) {
	c.Redirect(302, "https://v.ncuos.com/?appCode="+global.Config.AppCode)
	c.Abort()
}
