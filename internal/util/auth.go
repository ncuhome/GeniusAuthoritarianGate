package util

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"time"
)

func GoGeniusLogin(c *gin.Context) {
	c.Redirect(302, "https://v.ncuos.com/?appCode="+global.Config.AppCode)
	c.Abort()
}

const (
	RefreshTokenCookieKey = "refreshToken"
	AccessTokenCookieKey  = "accessToken"
)

func SetRefreshToken(c *gin.Context, token string) {
	c.SetCookie(RefreshTokenCookieKey, token, int((time.Duration(global.Config.LoginValidate) * time.Hour * 24).Seconds()), "", "", true, true)
}

func SetAccessToken(c *gin.Context, token string) {
	c.SetCookie(AccessTokenCookieKey, token, int((time.Minute * 5).Seconds()), "", "", true, true)
}
