package main

/**
 * @Description 启动文件
 * @Author Mikael
 * @Email 13247629622@163.com
 * @Date 2021/1/18 12:05
 * @Version 1.0
 **/

import (
	"flag"
	"fmt"
	"novel/job/internal/config"
	"novel/job/internal/handler"
	"novel/job/internal/svc"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
)

// 定义配置文件的位置
var configFile = flag.String("f", "etc/job.yaml", "the config file")

func main() {
	// 参数解析
	flag.Parse()

	// 加载配置
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 加载context
	ctx := svc.NewServiceContext(c)

	group := service.NewServiceGroup()
	// 注册定时任务
	handler.RegisterJob(ctx, group)

	// 信号处理
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-ch
		logx.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Printf("stop group")
			group.Stop()
			logx.Info("job exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
