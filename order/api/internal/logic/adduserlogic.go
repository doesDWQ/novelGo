package logic

import (
	"context"
	"novel/user/rpc/userclient"

	"novel/order/api/internal/svc"
	"novel/order/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddUserLogic {
	return AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req types.AddUserReq) (*types.AddUserReply, error) {
	// todo: add your logic here and delete this line

	_, err := l.svcCtx.UserRpc.AddUser(l.ctx, &userclient.AddUserRequest{
		Name:   "dwq",
		Gender: "1",
		Email:  "235@qq.com",
	})

	return &types.AddUserReply{}, err
}
