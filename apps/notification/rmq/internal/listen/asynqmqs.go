package listen

import (
	"Anitale/apps/notification/rmq/internal/config"
	"Anitale/apps/notification/rmq/internal/svc"
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/service"
)

type AsynqService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAsynqService(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqService {
	return &AsynqService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (r *AsynqService) Start() {
	srv := asynq.NewServer(asynq.RedisClientOpt{Addr: r.svcCtx.Config.Redis.Host, Password: r.svcCtx.Config.Redis.Pass},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		})

	mux := asynq.NewServeMux()
	mux.HandleFunc("", func(ctx context.Context, t *asynq.Task) error { return nil })

	if err := srv.Run(mux); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
	fmt.Println("asynq message server queue start...")
}

func (r *AsynqService) Stop() {
	fmt.Println("asynq message server queue stop...")
}

func AsynqMqs(c config.Config, ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		NewAsynqService(ctx, svcCtx),
	}
}
