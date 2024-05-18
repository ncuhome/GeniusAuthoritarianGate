package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/GeniusAuthoritarianClient/rpc/appProto"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/pkg/ga"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func Auth() gin.HandlerFunc {
	parser, err := ga.Rpc.NewJwtParser()
	if err != nil {
		log.Fatalln("create GeniusAuth jwt parser failed:", err)
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
			log.Warnln("get access cookie failed:", err)
		} else {
			_, valid, err := parser.ParseAccessToken(accessToken)
			if err != nil || !valid {
				log.Warnln("validate accessToken failed:", err)
			} else {
				return
			}
		}

		// Refresh accessToken
		refreshToken, err := util.GetRefreshToken(c)
		if err != nil || refreshToken == "" {
			log.Warnln("get refresh cookie failed:", err)
		} else {
			result, err := ga.Rpc.RefreshToken(context.Background(), &appProto.RefreshTokenRequest{
				Token: refreshToken,
			})
			if err != nil {
				if status.Code(err) != codes.Unauthenticated {
					log.Errorln("refresh access token failed:", err)
				}
			} else {
				err = util.SetAccessToken(c, result.AccessToken)
				if err != nil {
					log.Errorln("set access token failed:", err)
				}
				return
			}
		}

		util.ClearCookie(c)
		util.GoGeniusLogin(c)
	}
}
