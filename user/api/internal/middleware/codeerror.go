package middleware

import "net/http"

type CodeErrorMiddleware struct {
}

func NewCodeErrorMiddleware() *CodeErrorMiddleware {
	return &CodeErrorMiddleware{}
}

func (m *CodeErrorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}