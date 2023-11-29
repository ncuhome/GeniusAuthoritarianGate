package models

type Config struct {
	Addr string
	// default 30s
	Timeout uint `config:"omitempty"`

	// default 7d
	LoginValidate uint   `config:"omitempty"`
	WhiteListPath string `config:"omitempty"`

	AppCode   string
	AppSecret string
}
