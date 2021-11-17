package logic

/**
 * @Description 生产者任务
 * @Author Mikael
 * @Email 13247629622@163.com
 * @Date 2021/1/18 12:05
 * @Version 1.0
 **/

import (
	"context"
	"novel/job/internal/svc"

	"github.com/tal-tech/go-zero/core/threading"

	"github.com/tal-tech/go-zero/core/logx"
)

type Producer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProducerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Producer {
	return &Producer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Producer) Start() {

	logx.Infof("start  Producer \n")
	threading.GoSafe(func() {
		//producer := dq.NewProducer([]dq.Beanstalk{
		//	{
		//		Endpoint: "localhost:11300",
		//		Tube:     "tube1",
		//	},
		//	{
		//		Endpoint: "localhost:11301",
		//		Tube:     "tube2",
		//	},
		//})
		// producer.At([]byte("hello"), time.Date())
		//for i := 1000; i < 1005; i++ {
		//	_, err := producer.Delay([]byte(strconv.Itoa(i)), time.Second*1)
		//	if err != nil {
		//		logx.Error(err)
		//	}
		//}
	})
}

func (l *Producer) Stop() {
	logx.Infof("stop Producer \n")
}
