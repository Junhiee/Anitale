package logic

import (
	"context"

	"Anitale/apps/user/rpc/internal/svc"
	"Anitale/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPreferencesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPreferencesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPreferencesLogic {
	return &GetUserPreferencesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户偏好设置
func (l *GetUserPreferencesLogic) GetUserPreferences(in *pb.GetUserPreferencesRequest) (*pb.GetUserPreferencesResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserPreferencesResponse{}, nil
}
