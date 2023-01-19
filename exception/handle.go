package exception

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CodeError) Error() string {
	return e.Msg
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

type StrError struct {
	value string
}

func (e StrError) Error() string {
	return e.value
}
