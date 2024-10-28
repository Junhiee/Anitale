package user

import (
	"context"

	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"
	"Anitale/apps/user/rpc/pb"
	"Anitale/pkg/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserProfileLogic) GetUserProfile(req *types.GetUserProfileReq) (resp *types.GetUserProfileResp, err error) {
	userProfie, err := l.svcCtx.UserRpc.GetUserProfile(l.ctx, &pb.GetUserProfileRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomCode(errx.SERVER_ERROR), "UserRpc: user_id:%d,err:%v", req.UserId, err)
	}

	resp = &types.GetUserProfileResp{
		Profile: types.UserProfile{
			UserId:    userProfie.Profile.UserId,
			FullName:  userProfie.Profile.FullName,
			Bio:       userProfie.Profile.Bio,
			AvatarUrl: userProfie.Profile.AvatarUrl,
			Birthday:  userProfie.Profile.Birthday,
			Gender:    userProfie.Profile.Gender,
			Loc:       userProfie.Profile.Loc,
		},
	}

	return
}
