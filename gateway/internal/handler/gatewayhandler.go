package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"novel/gateway/internal/logic"
	"novel/gateway/internal/svc"
	"novel/gateway/internal/types"
)

func GatewayHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGatewayLogic(r.Context(), ctx)
		resp, err := l.Gateway(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
