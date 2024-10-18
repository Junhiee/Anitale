package logic

import (
	"context"
	"math"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"Anitale/apps/stats/rpc/internal/svc"
	"Anitale/apps/stats/rpc/pb"
)

type SortByHotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSortByHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SortByHotLogic {
	return &SortByHotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type AnimeHotScore struct {
	AnimeId int64
	Score   float64
}

// TODO 按热度排序, 返回排序完成的 anime 列表
func (l *SortByHotLogic) SortByHot(in *pb.SortByHotReq) (*pb.SortByHotResp, error) {

	return &pb.SortByHotResp{}, nil
}

// 热力值
func HotScore(stats *pb.Stats) float64 {
	// 权重定义
	alpha := 1.0 // 点击量权重
	beta := 0.5  // 点赞权重
	gamma := 0.3 // 评论权重
	delta := 0.7 // 分享权重

	// 时间衰减因子 (对数衰减)
	hoursSincePublished := time.Since(stats.LastUpdated.AsTime()).Hours()
	timeFactor := 1 / math.Log2(hoursSincePublished+2)

	// 热度分数计算
	score := (alpha*float64(stats.ViewCount) +
		beta*float64(stats.LikeCount) +
		gamma*float64(stats.CommentCount) +
		delta*float64(stats.ShareCount)) * timeFactor

	return score
}
