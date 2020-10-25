package shared

const (
	defaultNotFoundCode  = 4004
	defaultGRPCErrorCode = 1001
)

type (
	GRPCError struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
)

func NewGRPCError(code int, msg string) *GRPCError {
	return &GRPCError{Code: code, Msg: msg}
}

func NewGRPCNotFound() error {
	return NewGRPCError(defaultNotFoundCode, "")
}

func NewDefaultGRPCError(msg string) error {
	return NewGRPCError(defaultGRPCErrorCode, msg)
}

func NewGRPCErrorFromError(e error) error {
	return NewDefaultGRPCError(e.Error())
}

func (e *GRPCError) Error() string {
	return e.Msg
}

func FromGRPC(e error) error {
	switch v := e.(type) {
	case *GRPCError:
		return NewCodeError(v.Code, v.Msg)
	default:
		return NewDefaultError(e.Error())
	}
}

func IsGRPCNotFound(e error) bool {
	switch v := e.(type) {
	case *GRPCError:
		return v.Code == defaultNotFoundCode
	default:
		return false
	}
}
