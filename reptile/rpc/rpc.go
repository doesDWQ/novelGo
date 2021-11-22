package main

import (
	"flag"
	"novel/projectUtil"
	"novel/reptile/rpc/util"
)

// var configFile = flag.String("f", "etc/rpc.yaml", "the config file")

var log = projectUtil.GetFileLog()

func main() {
	flag.Parse()

	// 加载配置文件
	// var c config.Config
	// conf.MustLoad(*configFile, &c)
	// ctx := svc.NewServiceContext(c)

	// logx.Info("hello")

	log.Info(GetMenu())

}

// GetMenu 获取到菜单
func GetMenu() (data string) {
	data, err := util.HttpGet("https://www.xbiquge.la/?hpprotid=2a8031d1")

	if err != nil {
		panic(err)
	}

	return
}
