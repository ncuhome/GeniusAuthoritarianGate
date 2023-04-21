package middlewares

import (
	"github.com/gin-gonic/gin"
	ga "github.com/ncuhome/GeniusAuthoritarianClient"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	"github.com/sirupsen/logrus"
)

func Auth(jwtKey string) gin.HandlerFunc {
	gaClient := ga.NewClient("v.ncuos.com", util.Http.Client)
	return func(c *gin.Context) {
		switch c.Request.URL.Path {
		case "/login":
			token, ok := c.GetQuery("token")
			if !ok {
				c.AbortWithStatus(403)
				return
			}
			gaRes, e := gaClient.VerifyToken(&ga.RequestVerifyToken{
				Token:  token,
				Groups: global.AllowGroups,
			})
			if e != nil {
				logrus.Errorln("GeniusAuth 身份校验异常:", e)
				c.AbortWithStatus(500)
				return
			} else if gaRes.Code != 0 {
				logrus.Errorln("GeniusAuth 身份校验失败:", gaRes.Msg)
				c.String(403, gaRes.Msg)
				return
			}
			c.Redirect(302, "/")
		default:
			/*token,e:=c.Cookie("token")
			if e!=nil {
				logrus.Warnln("无法处理 cookie:",e)
				c.AbortWithStatus(400)
				return
			}

			*/
		}
	}
}
