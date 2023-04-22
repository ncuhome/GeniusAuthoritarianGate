package util

import (
	"github.com/gin-gonic/gin"
	"net/url"
)

func GoGeniusLogin(c *gin.Context) {
	c.Redirect(302, "https://v.ncuos.com/?target="+url.QueryEscape("https://"+c.Request.Host+"/login"))
	c.Abort()
}
