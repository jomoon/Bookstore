package handler

import (
	"net/http"

	"Bookstore/api/internal/logic"
	"Bookstore/api/internal/svc"
	"Bookstore/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func LoginerHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginerLogic(r.Context(), ctx)
		resp, err := l.Loginer(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
