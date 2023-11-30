package callback

const (
	cErrLoginNeeded uint8 = iota + 1
	cErrLoginFailed
	cErrUnexpected
)

var (
	ErrLoginNeeded = &Msg{
		Code:       cErrLoginNeeded,
		Msg:        "请重新登录",
		HttpStatus: 401,
	}
	ErrLoginFailed = &Msg{
		Code:       cErrLoginFailed,
		Msg:        "登录出错，请重新登录",
		HttpStatus: 500,
	}
	ErrUnexpected = &Msg{
		Code:       cErrUnexpected,
		Msg:        "发生预期外错误，请反馈开发者",
		HttpStatus: 500,
	}
)
