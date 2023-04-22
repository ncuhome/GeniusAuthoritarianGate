package models

type Config struct {
	Addr string
	// default 30s
	Timeout uint `config:"omitempty"`
	// default all
	Groups string `config:"omitempty"`
	JwtKey string `config:"omitempty"`
	// default 7d
	LoginValidate uint `config:"omitempty"`
}
