package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	refreshTokenRpc "github.com/ncuhome/GeniusAuthoritarianRefreshTokenRpc"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Auth() (gin.HandlerFunc, error) {
	rpcClient, err := refreshTokenRpc.NewRpc("v.ncuos.com:443", &refreshTokenRpc.Config{
		AppCode:   global.Config.AppCode,
		AppSecret: global.Config.AppSecret,
	})
	if err != nil {
		return nil, err
	}
	return func(c *gin.Context) {
		if c.FullPath() != "" {
			// 跳过 api
			return
		}

		for _, whiteListPath := range global.WhiteListPath {
			if c.Request.URL.Path == whiteListPath {
				return
			}
		}

		accessToken, err := util.GetAccessToken(c)
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

		refreshToken, err := util.GetRefreshToken(c)
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
				err = util.SetAccessToken(c, result.AccessToken)
				if err != nil {
					log.Errorln("设置 access token 失败:", err)
					c.String(500, "编码 access token 失败")
					return
				}
				return
			}
		}

		util.GoGeniusLogin(c)
	}, nil
}
