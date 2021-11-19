package logic

import (
	"context"

	"novel/websocket/server/internal/svc"
	"novel/websocket/server/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServerLogic {
	return ServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServerLogic) Server(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
