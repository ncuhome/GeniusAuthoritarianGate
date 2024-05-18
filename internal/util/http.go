package util

import (
	"github.com/Mmx233/tool"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"net"
	"net/http"
	"time"
)

var Http *tool.Http

func init() {
	timeout := time.Second * time.Duration(global.Config.Timeout)

	Http = tool.NewHttpTool(tool.GenHttpClient(&tool.HttpClientOptions{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: timeout,
			}).DialContext,
		},
		Timeout: timeout,
	}))
}
