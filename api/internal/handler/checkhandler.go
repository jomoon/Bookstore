package handler

import (
	"net/http"

	"Bookstore/api/internal/logic"
	"Bookstore/api/internal/svc"
	"Bookstore/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func CheckHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCheckLogic(r.Context(), ctx)
		resp, err := l.Check(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
