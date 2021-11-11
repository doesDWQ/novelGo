package logic

import (
	"context"
	"novel/user/rpc/internal/svc"
	"novel/user/rpc/model"
	"novel/user/rpc/user"
	"strconv"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {

	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.AddUserRequest) (*user.AddUserResponse, error) {
	// todo: add your logic here and delete this line

	userId := "1"

	gender, err := strconv.Atoi(in.Gender)
	if err != nil {
		return nil, err
	}

	// 手动代码开始
	_, err = l.svcCtx.Model.Insert(model.User{
		Id:     userId,
		Name:   in.Name,
		Gender: int64(gender),
		Email:  in.Email,
	})

	return &user.AddUserResponse{
		Id:     userId,
		Status: "200",
	}, err
}
