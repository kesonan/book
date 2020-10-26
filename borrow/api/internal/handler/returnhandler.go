package handler

import (
	"net/http"

	"book/borrow/api/internal/logic"
	"book/borrow/api/internal/svc"
	"book/borrow/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func returnHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReturnReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewReturnLogic(r.Context(), ctx)
		err := l.Return(r.Header.Get("x-user-id"), req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
