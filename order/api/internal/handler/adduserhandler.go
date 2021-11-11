package handler

import (
	"net/http"

	"novel/order/api/internal/logic"
	"novel/order/api/internal/svc"
	"novel/order/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func addUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddUserLogic(r.Context(), ctx)
		resp, err := l.AddUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
