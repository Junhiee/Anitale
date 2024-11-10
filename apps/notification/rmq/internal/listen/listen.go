package listen

import (
	"Anitale/apps/notification/rmq/internal/config"
	"Anitale/apps/notification/rmq/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/service"
)

func Mqs(c config.Config) []service.Service {
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()

	var service []service.Service

	service = append(service, AsynqMqs(c, ctx, svcCtx)...)
	return service
}
