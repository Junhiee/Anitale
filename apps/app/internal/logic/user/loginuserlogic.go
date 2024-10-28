package user

import (
	"context"

	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"
	"Anitale/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginUserLogic) LoginUser(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	item, err := l.svcCtx.UserRpc.LoginUser(l.ctx, &pb.LoginUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UserLoginResp{
		AccessToken: item.AccessToken,
		Expire:      item.Expire,
	}

	return
}
