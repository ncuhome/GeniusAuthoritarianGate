package middlewares

import (
	"github.com/gin-gonic/gin"
	ga "github.com/ncuhome/GeniusAuthoritarianClient"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	refreshTokenRpc "github.com/ncuhome/GeniusAuthoritarianRefreshTokenRpc"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func Auth() (gin.HandlerFunc, error) {
	gaClient := ga.NewClient(
		"v.ncuos.com",
		global.Config.AppCode, global.Config.AppSecret,
		util.Http.Client,
	)
	rpcClient, err := refreshTokenRpc.NewRpc("v.ncuos.com:443", &refreshTokenRpc.Config{
		AppCode:   global.Config.AppCode,
		AppSecret: global.Config.AppSecret,
	})
	if err != nil {
		return nil, err
	}
	return func(c *gin.Context) {
		switch c.Request.URL.Path {
		case "/login":
			token, ok := c.GetQuery("token")
			if !ok {
				tokenCookie, err := c.Cookie(util.RefreshTokenCookieKey)
				if err != nil || tokenCookie == "" {
					util.GoGeniusLogin(c)
				} else {
					c.Redirect(302, "/")
				}
				return
			}
			gaRes, err := gaClient.VerifyToken(&ga.RequestVerifyToken{
				Token:     token,
				ClientIp:  c.ClientIP(),
				GrantType: "refresh_token",
				Valid:     int64((time.Duration(global.Config.LoginValidate) * time.Hour * 24).Seconds()),
			})
			if err != nil {
				log.Errorln("GeniusAuth 身份校验异常:", err)
				c.String(500, "身份校验异常")
				return
			} else if gaRes.Code != 0 {
				log.Errorln("GeniusAuth 身份校验失败:", gaRes.Msg)
				c.String(403, gaRes.Msg)
				return
			}

			// refreshToken 不能一直发送到服务端，但是此处没有前端不好写，先临时这样处理
			util.SetRefreshToken(c, gaRes.Data.RefreshToken)
			util.SetAccessToken(c, gaRes.Data.AccessToken)
			c.Redirect(302, "/")
		default:
			for _, whiteListPath := range global.WhiteListPath {
				if c.Request.URL.Path == whiteListPath {
					return
				}
			}

			accessToken, err := c.Cookie(util.AccessTokenCookieKey)
			if err != nil || accessToken == "" {
				log.Warnln("无法获取 access cookie:", err)
			} else {
				_, err = rpcClient.VerifyAccessToken(context.Background(), accessToken)
				if err != nil {
					log.Warnln("验证 accessToken 失败:", err)
				} else {
					return
				}
			}

			// Refresh accessToken
			refreshToken, err := c.Cookie(util.RefreshTokenCookieKey)
			if err != nil || accessToken == "" {
				log.Warnln("无法获取 refresh cookie:", err)
			} else {
				result, err := rpcClient.RefreshToken(context.Background(), refreshToken)
				if err != nil {
					if status.Code(err) != codes.Unauthenticated {
						log.Errorln("刷新 token 异常:", err)
						c.String(500, "刷新 token 异常")
						return
					}
				} else {
					util.SetAccessToken(c, result.AccessToken)
					return
				}
			}

			util.GoGeniusLogin(c)
		}
	}, nil
}
