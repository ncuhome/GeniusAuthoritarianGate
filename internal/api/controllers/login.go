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
	token, ok := c.GetQuery("token")
	if !ok || token == "" {
		tokenCookie, err := util.GetRefreshToken(c)
		if err != nil || tokenCookie == "" {
			callback.Error(c, callback.ErrLoginNeeded, err)
		} else {
			callback.Default(c)
		}
		return
	}

	gaRes, err := ga.Client.VerifyToken(&geniusAuth.RequestVerifyToken{
		Token:     token,
		ClientIp:  c.ClientIP(),
		GrantType: "refresh_token",
		Valid:     int64((time.Duration(global.Config.LoginValidate) * time.Hour * 24).Seconds()),
	})
	if err != nil {
		callback.ErrorWithTip(c, callback.ErrLoginFailed, "身份校验异常", err)
		return
	} else if gaRes.Code != 0 {
		callback.ErrorWithTip(c, callback.ErrLoginFailed, gaRes.Msg, err)
		return
	}

	err = util.SetRefreshToken(c, gaRes.Data.RefreshToken)
	if err != nil {
		callback.Error(c, callback.ErrUnexpected, err)
		return
	}

	err = util.SetAccessToken(c, gaRes.Data.AccessToken)
	if err != nil {
		callback.Error(c, callback.ErrUnexpected, err)
		return
	}

	callback.Default(c)
}
