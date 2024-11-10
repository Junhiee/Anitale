package user

import (
	"context"

	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户订阅接口
func NewUserSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSubscribeLogic {
	return &UserSubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSubscribeLogic) UserSubscribe(req *types.UserSubscribeReq) (resp *types.UserSubscribeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
