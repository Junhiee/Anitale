package logic

import (
	"context"

	"Anitale/apps/user/rpc/internal/svc"
	"Anitale/apps/user/rpc/model"
	"Anitale/apps/user/rpc/pb"
	"Anitale/pkg/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserSubscribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSubscribeLogic {
	return &UserSubscribeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户订阅
func (l *UserSubscribeLogic) UserSubscribe(in *pb.UserSubscribeRequest) (*pb.UserSubscribeResponse, error) {
	// 通知偏好请求
	// preference, err := l.svcCtx.UserPreferencesModel.FindOne(l.ctx, in.UserId)
	// if err != nil {
	// 	return nil, errors.Wrapf(errx.NewCustomCode(errx.SERVER_ERROR), "user_id:%d,err:%v", in.UserId, err)
	// }

	err := l.svcCtx.UserSubscriptionsModel.Insert(l.ctx, nil, &model.UserSubscriptions{
		UserId:  in.UserId,
		AnimeId: in.AnimeId,
	})
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomCode(errx.DB_ERROR), "user_id:%d,anime_id:%d,err:%v", in.UserId, in.AnimeId, err)
	}

	return &pb.UserSubscribeResponse{
		Subscription: &pb.UserSubscriptions{
			UserId:  in.UserId,
			AnimeId: in.AnimeId,
			Status:  "active",
		},
	}, nil
}
