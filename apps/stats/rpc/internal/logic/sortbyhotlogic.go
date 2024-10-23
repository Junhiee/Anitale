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

func (l *SortByHotLogic) SortByHot(in *pb.SortByHotReq) (*pb.SortByHotResp, error) {
	// 如果redis中没有排行榜的数据, 则初始化排行榜
	// 如果有数据并且命中了缓存，则返回查询的结果值
	// 如果没有命中缓存，则从数据库中查询并计算热力值，并写入缓存

	return &pb.SortByHotResp{}, nil
}

// 热力值
// select *, (view_count * 1.0 + like_count * 0.5 + comment_count * 0.3 + share_count * 0.7) / POWER(TIMESTAMPDIFF(HOUR, last_updated, NOW()) + 2, 1.5) AS hot_score from stats order by hot_score desc;

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
