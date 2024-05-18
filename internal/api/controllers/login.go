package controllers

import (
	"github.com/gin-gonic/gin"
	geniusAuth "github.com/ncuhome/GeniusAuthoritarianClient"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/api/callback"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/pkg/ga"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	"time"
)

func Login(c *gin.Context) {
	token, ok := c.GetPostForm("token")
	if !ok || token == "" {
		callback.Error(c, callback.ErrLoginNeeded)
		return
	}

	result, err := ga.Client.VerifyToken(&geniusAuth.RequestVerifyToken{
		Token:     token,
		ClientIp:  c.ClientIP(),
		GrantType: "refresh_token",
		Valid:     int64((time.Duration(global.Config.LoginValidate) * time.Hour * 24).Seconds()),
	})
	if err != nil {
		callback.ErrorWithTip(c, callback.ErrLoginFailed, "GeniusAuth login request verify failed", err)
		return
	}

	err = util.SetRefreshToken(c, result.RefreshToken)
	if err != nil {
		callback.Error(c, callback.ErrUnexpected, err)
		return
	}

	err = util.SetAccessToken(c, result.AccessToken)
	if err != nil {
		callback.Error(c, callback.ErrUnexpected, err)
		return
	}

	callback.Default(c)
}
