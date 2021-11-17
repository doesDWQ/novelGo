package logic

/**
 * @Description 消费者任务
 * @Author Mikael
 * @Email 13247629622@163.com
 * @Date 2021/1/18 12:05
 * @Version 1.0
 **/

import (
	"context"
	"novel/job/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type Consumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Consumer {
	return &Consumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Consumer) Start() {
	logx.Infof("start consumer \n")

	//threading.GoSafe(func() {
	//	l.svcCtx.Consumer.Consume(func(body []byte) {
	//		logx.Infof("consumer job  %s \n" ,string(body))
	//	})
	//})
}

func (l *Consumer) Stop() {
	logx.Infof("stop consumer \n")
}
