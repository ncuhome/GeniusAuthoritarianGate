package middlewares

import (
	"github.com/gin-gonic/gin"
	ga "github.com/ncuhome/GeniusAuthoritarianClient"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/pkg/jwt"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/util"
	log "github.com/sirupsen/logrus"
	"time"
)

func Auth(jwtKey string) gin.HandlerFunc {
	gaClient := ga.NewClient("v.ncuos.com", util.Http.Client)
	jwtHandler := jwt.New([]byte(jwtKey), time.Hour*24*time.Duration(global.Config.LoginValidate))
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
				log.Errorln("GeniusAuth 身份校验异常:", e)
				c.AbortWithStatus(500)
				return
			} else if gaRes.Code != 0 {
				log.Errorln("GeniusAuth 身份校验失败:", gaRes.Msg)
				c.String(403, gaRes.Msg)
				return
			}

			token, e = jwtHandler.NewToken()
			if e != nil {
				log.Errorln("生成 token 失败:", e)
				c.AbortWithStatus(500)
				return
			}
			c.SetCookie("token", token, int(jwtHandler.Validate.Seconds()), "", "", true, true)
			c.Redirect(302, "/")
		default:
			token, e := c.Cookie("token")
			if e != nil || token == "" {
				log.Warnln("无法处理 cookie:", e)
				util.GoGeniusLogin(c)
				return
			}

			valid, e := jwtHandler.VerifyToken(token)
			if e != nil || !valid {
				log.Warnln("身份校验失败:", e)
				util.GoGeniusLogin(c)
				return
			}
		}
	}
}
