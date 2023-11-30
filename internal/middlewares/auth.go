package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	refreshTokenRpc "github.com/ncuhome/GeniusAuthoritarianRefreshTokenRpc"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func Auth() gin.HandlerFunc {
	rpcClient, err := refreshTokenRpc.NewRpc(fmt.Sprintf("%s:443", global.Config.GeniusAuthHost), &refreshTokenRpc.Config{
		AppCode:   global.Config.AppCode,
		AppSecret: global.Config.AppSecret,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/login/") {
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
		if err != nil || refreshToken == "" {
			log.Warnln("无法获取 refresh cookie:", err)
		} else {
			result, err := rpcClient.RefreshToken(context.Background(), refreshToken)
			if err != nil {
				if status.Code(err) != codes.Unauthenticated {
					log.Errorln("刷新 access token 异常:", err)
				}
			} else {
				err = util.SetAccessToken(c, result.AccessToken)
				if err != nil {
					log.Errorln("设置 access token 失败:", err)
				}
				return
			}
		}

		util.ClearCookie(c)
		util.GoGeniusLogin(c)
	}
}
