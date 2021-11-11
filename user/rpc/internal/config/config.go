package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	DataSource string          // 手动代码 加载数据库配置
	Cache      cache.CacheConf // 手动代码 加载模型缓存配置
}
