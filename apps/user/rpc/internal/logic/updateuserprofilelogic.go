package logic

import (
	"Anitale/apps/user/rpc/internal/svc"
	"Anitale/apps/user/rpc/model"
	"Anitale/apps/user/rpc/pb"
	"Anitale/pkg/e"
	"Anitale/pkg/util"
	"context"

	"github.com/pkg/errors"
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

	var resp = &pb.UpdateUserProfileResponse{
		Success: true,
		Message: "success",
		Profile: &pb.UserProfile{
			UserId:    in.UserId,
			FullName:  in.FullName,
			Bio:       in.Bio,
			AvatarUrl: in.AvatarUrl,
			Birthday:  in.Birthday,
			Gender:    in.Gender,
			Loc:       in.Loc,
		},
	}

	err := l.svcCtx.UserProfilesModel.Update(l.ctx, nil, &model.UserProfiles{
		UserId:    in.UserId,
		FullName:  util.NewNullString(in.FullName),
		Bio:       util.NewNullString(in.Bio),
		AvatarUrl: util.NewNullString(in.AvatarUrl),
		Birthday:  util.NewNullTime(in.Birthday.AsTime()),
		Gender:    util.NewNullString(in.Gender),
		Loc:       util.NewNullString(in.Loc),
	})

	if err != nil {
		return nil, errors.Wrapf(e.NewCustomCode(e.DB_ERROR), "UserProfilesModel.Update user_id:%d,err:%v", in.UserId, err)
	}

	return resp, nil
}
