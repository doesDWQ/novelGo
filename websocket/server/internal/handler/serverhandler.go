package handler

import (
	"net/http"
	"novel/websocket/server/internal/svc"
	"novel/websocket/server/internal/types"
	"novel/websocket/server/internal/ws"

	"github.com/tal-tech/go-zero/rest/httpx"
)

var WsHub = ws.NewHub()

func ServerHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request

		// 解析request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		ws.ServeWs(WsHub, w, r, ctx)
	}
}
