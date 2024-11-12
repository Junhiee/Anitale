package logic

import (
	"context"

	"Anitale/apps/user/rpc/internal/svc"
	"Anitale/apps/user/rpc/model"
	"Anitale/apps/user/rpc/pb"
	"Anitale/pkg/e"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户个人信息
func (l *GetUserProfileLogic) GetUserProfile(in *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	var resp = &pb.GetUserProfileResponse{}
	// 查找用户
	userProfile, err := l.svcCtx.UserProfilesModel.FindOne(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(e.NewCustomCode(e.DB_ERROR), "user_id:%d,err:%v", in.UserId, err)
	}

	// 用户不存在
	if userProfile == nil {
		return nil, errors.Wrapf(e.NewCustomCode(e.USER_NOT_FOUND_ERROR), "user_id:%d, err:%v", in.UserId, err)
	}

	resp.Profile = &pb.UserProfile{
		UserId:    userProfile.UserId,
		FullName:  userProfile.FullName.String,
		Bio:       userProfile.Bio.String,
		AvatarUrl: userProfile.AvatarUrl.String,
		Birthday:  timestamppb.New(userProfile.Birthday.Time),
		Gender:    userProfile.Gender.String,
		Loc:       userProfile.Loc.String,
	}

	return resp, nil
}
