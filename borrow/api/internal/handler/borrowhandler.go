package handler

import (
	"net/http"

	"book/borrow/api/internal/logic"
	"book/borrow/api/internal/svc"
	"book/borrow/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func borrowHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BorrowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewBorrowLogic(r.Context(), ctx)
		err := l.Borrow(r.Header.Get("x-user-id"), req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
