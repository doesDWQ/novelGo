package handler

/**
 * @Description 注册job
 * @Author Mikael
 * @Email 13247629622@163.com
 * @Date 2021/1/18 12:05
 * @Version 1.0
 **/

import (
	"context"
	"novel/job/internal/logic"
	"novel/job/internal/svc"

	"github.com/tal-tech/go-zero/core/service"
)

func RegisterJob(serverCtx *svc.ServiceContext, group *service.ServiceGroup) {

	group.Add(logic.NewProducerLogic(context.Background(), serverCtx))
	group.Add(logic.NewConsumerLogic(context.Background(), serverCtx))

	// 启动定时任务
	group.Start()
}
