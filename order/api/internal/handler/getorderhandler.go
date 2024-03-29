package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"novel/order/api/internal/logic"
	"novel/order/api/internal/svc"
	"novel/order/api/internal/types"
)

func getOrderHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetOrderLogic(r.Context(), ctx)
		resp, err := l.GetOrder(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
