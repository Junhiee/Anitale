package logic

import (
	"context"

	"Anitale/apps/user/rpc/internal/svc"
	"Anitale/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPreferencesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPreferencesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPreferencesLogic {
	return &UpdateUserPreferencesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户偏好设置
func (l *UpdateUserPreferencesLogic) UpdateUserPreferences(in *pb.UpdateUserPreferencesRequest) (*pb.UpdateUserPreferencesResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserPreferencesResponse{}, nil
}
