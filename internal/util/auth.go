package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/pkg/aes"
	"time"
)

func GoGeniusLogin(c *gin.Context) {
	c.Redirect(302, fmt.Sprintf("https://%s/?appCode=%s", global.Config.GeniusAuthHost, global.Config.AppCode))
	c.Abort()
}

const (
	RefreshTokenCookieKey = "refreshToken"
	AccessTokenCookieKey  = "accessToken"
)

func SetRefreshToken(c *gin.Context, token string) error {
	encryptedToken, err := aes.EncryptString(token)
	if err != nil {
		return err
	}
	c.SetCookie(RefreshTokenCookieKey, encryptedToken, int((time.Duration(global.Config.LoginValidate) * time.Hour * 24).Seconds()), "", "", true, true)
	return nil
}

func GetRefreshToken(c *gin.Context) (string, error) {
	tokenStr, err := c.Cookie(RefreshTokenCookieKey)
	if err != nil {
		return "", err
	}
	return aes.DecryptString(tokenStr)
}

func SetAccessToken(c *gin.Context, token string) error {
	encryptedToken, err := aes.EncryptString(token)
	if err != nil {
		return err
	}
	c.SetCookie(AccessTokenCookieKey, encryptedToken, int((time.Minute * 5).Seconds()), "", "", true, true)
	return nil
}

func GetAccessToken(c *gin.Context) (string, error) {
	tokenStr, err := c.Cookie(AccessTokenCookieKey)
	if err != nil {
		return "", err
	}
	return aes.DecryptString(tokenStr)
}

func ClearCookie(c *gin.Context) {
	c.SetCookie(RefreshTokenCookieKey, "", -1, "", "", true, true)
	c.SetCookie(AccessTokenCookieKey, "", -1, "", "", true, true)
}
