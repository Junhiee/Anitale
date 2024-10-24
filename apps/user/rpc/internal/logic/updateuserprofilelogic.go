package logic

import (
	"context"

	"Anitale/apps/user/rpc/internal/svc"
	"Anitale/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户个人信息
func (l *UpdateUserProfileLogic) UpdateUserProfile(in *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserProfileResponse{}, nil
}
