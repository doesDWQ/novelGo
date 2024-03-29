package svc

import (
	"novel/user/rpc/internal/config"
	"novel/user/rpc/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	Model model.UserModel // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
	}
}
