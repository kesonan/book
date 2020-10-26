package shared

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	defaultNotFoundCode     = 4004
	defaultGRPCErrorCode    = 1001
	defaultErrorPlaceholder = ""
)

func NewGRPCCodeError(code int, msg string) error {
	st := status.New(codes.Code(code), msg)
	return st.Err()
}

func NewGRPCNotFound() error {
	st := status.New(codes.Code(defaultNotFoundCode), defaultErrorPlaceholder)
	return st.Err()
}

func NewDefaultGRPCError(msg string) error {
	return NewGRPCCodeError(defaultGRPCErrorCode, msg)
}

func NewGRPCErrorFromError(e error) error {
	return NewDefaultGRPCError(e.Error())
}

func FromGRPC(e error) error {
	st := status.Convert(e)
	switch st.Code() {
	case codes.OK:
		return st.Err()
	default:
		return NewCodeError(int(st.Code()), st.Message())
	}
}

func IsGRPCNotFound(e error) bool {
	st := status.Convert(e)
	switch st.Code() {
	case codes.OK:
		return false
	default:
		return int(st.Code()) == defaultNotFoundCode
	}
}
