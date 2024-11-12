package logic

import (
	"Anitale/apps/user/rpc/model"
	"Anitale/pkg/e"
	"context"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"Anitale/apps/user/rpc/internal/svc"
	"Anitale/apps/user/rpc/pb"

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
//func (l *UserSubscribeLogic) UserSubscribe(in *pb.UserSubscribeRequest) (*pb.UserSubscribeResponse, error) {
//	// todo: add your logic here and delete this line
//
//	return &pb.UserSubscribeResponse{}, nil
//}

// UserSubscribe 用户订阅
func (l *UserSubscribeLogic) UserSubscribe(in *pb.UserSubscribeRequest) (*pb.UserSubscribeResponse, error) {
	// TODO 通知偏好
	// 用户是否通过邮件通知
	// 用户是否通过短信通知

	var (
		subscribedAt           = time.Now()
		notificationPreference = "web"    // 默认通过web平台通知消息
		status                 = "active" // 状态：已订阅
		resp                   = &pb.UserSubscribeResponse{
			UserId:                 in.UserId,
			AnimeId:                in.AnimeId,
			NotificationPreference: notificationPreference,
			Status:                 status,
			SubscribedAt:           timestamppb.New(subscribedAt),
		}
	)
	// 用户是否已经订阅
	isSubscribed, err := l.svcCtx.UserSubscriptionsModel.IsSubscribed(l.ctx, in.UserId, in.AnimeId)
	if err != nil {
		return nil, errors.Wrapf(e.NewCustomCode(e.DB_ERROR), "UserSubscriptionsModel.IsSubscribed user_id:%d,anime_id:%d,err:%v", in.UserId, in.AnimeId, err)
	}
	// 如果没有订阅，则订阅
	if !isSubscribed {
		err = l.svcCtx.UserSubscriptionsModel.Insert(l.ctx, nil, &model.UserSubscriptions{
			UserId:                 in.UserId,
			AnimeId:                in.AnimeId,
			SubscribedAt:           subscribedAt,
			NotificationPreference: notificationPreference,
			Status:                 status,
		})
		if err != nil {
			return nil, errors.Wrapf(e.NewCustomCode(e.DB_ERROR), "UserSubscriptionsModel.Insert user_id:%d,anime_id:%d,err:%v", in.UserId, in.AnimeId, err)
		}
	}

	// 已经订阅的，再次订阅就会取消订阅
	if isSubscribed {
		err = l.svcCtx.UserSubscriptionsModel.DeleteByAnimeIdAndUserId(l.ctx, in.UserId, in.AnimeId)
		if err != nil {
			return nil, errors.Wrapf(e.NewCustomCode(e.DB_ERROR), "UserSubscriptionsModel.IsSubscribed user_id:%d,anime_id:%d,err:%v", in.UserId, in.AnimeId, err)
		}

		// 已经取消订阅的状态
		resp.Status = "inactive"

	}

	return resp, nil
}
