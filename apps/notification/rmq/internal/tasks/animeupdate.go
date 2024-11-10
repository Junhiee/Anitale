package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"Anitale/apps/anime/rpc/model"
	"Anitale/apps/notification/rmq/internal/svc"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

type AsynqTask struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAsynqTask(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqTask {
	return &AsynqTask{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type AnimeUpdateType struct {
	model.AnimeUpdates
}

// TODO: 邮件推送 | 定时任务 | 延迟任务
func (r *AsynqTask) AnimeUpdateHandler(ctx context.Context, t *asynq.Task) error {
	var p AnimeUpdateType
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	// anime_id := p.AnimeId
	// user rpc -> 通过查询 `user_subscriptions` 表，获取订阅该动画的所有 `user_id`
	// 发送通知
	// 推送后通知记录插入 notifications 表

	return nil
}
