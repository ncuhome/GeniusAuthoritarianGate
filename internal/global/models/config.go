package models

type Config struct {
	Addr    string
	Timeout uint   `config:"omitempty"`
	Groups  string `config:"omitempty"`
}
