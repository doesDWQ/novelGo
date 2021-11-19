package handler

import (
	"net/http"
	"novel/websocket/server/internal/svc"
	"novel/websocket/server/internal/types"
	"novel/websocket/server/internal/ws"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ServerHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		wsHub := ws.NewHub()
		go wsHub.Run()
		ws.ServeWs(wsHub, w, r, ctx)
	}
}
