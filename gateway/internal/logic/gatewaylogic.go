package logic

import (
	"context"

	"novel/gateway/internal/svc"
	"novel/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GatewayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGatewayLogic(ctx context.Context, svcCtx *svc.ServiceContext) GatewayLogic {
	return GatewayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GatewayLogic) Gateway(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
