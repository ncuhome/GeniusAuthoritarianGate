package util

import (
	"github.com/Mmx233/tool"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"time"
)

var Http *tool.Http

func init() {
	defaultTimeout := time.Second * time.Duration(global.Config.Timeout)

	Http = tool.NewHttpTool(tool.GenHttpClient(&tool.HttpClientOptions{
		Transport: tool.GenHttpTransport(&tool.HttpTransportOptions{
			Timeout: defaultTimeout,
		}),
		Timeout: defaultTimeout,
	}))
}
