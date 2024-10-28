package user

import (
	"context"

	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"
	"Anitale/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterUserLogic) RegisterUser(req *types.RegisterUserReq) (resp *types.RegisterUserResp, err error) {
	resp = &types.RegisterUserResp{}

	items, err := l.svcCtx.UserRpc.RegisterUser(l.ctx, &pb.RegisterUserRequest{
		Email:    req.Email,
		Password: req.Password,
		Username: req.UserName,
	})
	if err != nil {
		return nil, err
	}


	resp.User = types.User{
		Id:         items.User.Id,
		Email:      req.Email,
		UserName:   req.UserName,
		IsActive:   items.User.IsActive,
		IsVerified: items.User.IsActive,
	}

	return
}
