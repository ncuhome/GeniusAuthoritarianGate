package util

import (
	"github.com/Mmx233/tool"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	"net/http"
	"time"
)

var HttpTransport *http.Transport

func init() {
	HttpTransport = tool.GenHttpTransport(&tool.HttpTransportOptions{
		Timeout: time.Second * time.Duration(global.Config.Timeout),
	})
}
