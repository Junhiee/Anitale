package logic

import (
	"Anitale/apps/anime/rpc/model"
	"Anitale/pkg/errx"
	"Anitale/pkg/util"
	"context"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math"
	"strconv"
	"time"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

type RankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RankLogic {
	return &RankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RankLogic) Rank(in *pb.RankReq) (*pb.AnimeListResp, error) {
	var err error
	var resp = &pb.AnimeListResp{}
	var items = []model.Anime{}

	cmds, _ := l.svcCtx.RedisClient.Pipelined(l.ctx, func(p redis.Pipeliner) error {
		for _, item := range items {
			err = p.ZScore(l.ctx, "cache:hot:rank", strconv.FormatInt(item.AnimeId, 10)).Err()
			if err != nil {
				return err
			}
		}
		return nil
	})

	for i, cmd := range cmds {
		// 如果该 anime_id 没有从redis中获得热力值
		if cmd.Err() == redis.Nil {
			// 从 hset 的 stats 集合中计算热力，并存储在 redis 的 zset 中
			hcmd := l.svcCtx.RedisClient.HGetAll(l.ctx, "cache:stats:"+strconv.FormatInt(items[i].AnimeId, 10))
			// 如果 hset 中也没有该 anime_id 的 stats 数据
			// 则从数据库获取并同步更新 redis 中的 stats 数据
			if len(hcmd.Val()) == 0 {
				stats, err := l.svcCtx.StatsModel.FindOneNoCacheCtx(l.ctx, items[i].AnimeId)
				if err != nil {
					return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "StatsModel FindOne Err: %v", err)
				}

				// stats struct to map
				m := map[string]interface{}{
					"view_count":    stats.ViewCount,
					"like_count":    stats.LikeCount,
					"comment_count": stats.CommentCount,
					"share_count":   stats.ShareCount,
					"updated_at":    stats.UpdatedAt,
				}
				score := CalculateHeatValue(*stats, time.Now())

				// 同步更新 redis 中的 stats 数据
				l.svcCtx.RedisClient.HSet(l.ctx, "cache:stats:"+strconv.FormatInt(items[i].AnimeId, 10), m)

				// 计算热力值，存储到 cache:hot:rank 中
				l.svcCtx.RedisClient.ZAdd(l.ctx, "cache:hot:rank", redis.Z{
					Score:  score,
					Member: strconv.FormatInt(items[i].AnimeId, 10),
				})

				// 计算热力值，存储到 cache:hot:temp 中
				l.svcCtx.RedisClient.ZAdd(l.ctx, "cache:hot:temp", redis.Z{
					Score:  score,
					Member: strconv.FormatInt(items[i].AnimeId, 10),
				})
			}

			// 如果有 stats 数据
			if len(hcmd.Val()) != 0 {

				// 计算热力值
				view_count, _ := strconv.Atoi(hcmd.Val()["view_count"])
				like_count, _ := strconv.Atoi(hcmd.Val()["like_count"])
				comment_count, _ := strconv.Atoi(hcmd.Val()["comment_count"])
				share_count, _ := strconv.Atoi(hcmd.Val()["share_count"])
				updated_at, _ := strconv.Atoi(hcmd.Val()["updated_at"])
				score := CalculateHeatValue(model.Stats{
					ViewCount:    int64(view_count),
					LikeCount:    int64(like_count),
					CommentCount: int64(comment_count),
					ShareCount:   int64(share_count),
					CreatedAt:    time.Unix(int64(updated_at), 0),
				}, time.Now())

				// 存入redis中
				l.svcCtx.RedisClient.ZAdd(l.ctx, "cache:hot:rank", redis.Z{
					Score:  score,
					Member: strconv.FormatInt(items[i].AnimeId, 10),
				})
				l.svcCtx.RedisClient.ZAdd(l.ctx, "cache:hot:temp", redis.Z{
					Score:  score,
					Member: strconv.FormatInt(items[i].AnimeId, 10),
				})
			}

		}

		// 如果有获得热力值，则存入到 cache:hot:temp 中
		if cmd.Err() != redis.Nil {
			l.svcCtx.RedisClient.ZAdd(l.ctx, "cache:hot:temp", redis.Z{
				Score:  cmd.(*redis.FloatCmd).Val(),
				Member: strconv.FormatInt(items[i].AnimeId, 10),
			})
		}

	}

	// 取出 cache:hot:temp 中的数据，按热力值排序并分页返回
	sorted, err := l.svcCtx.RedisClient.ZRevRange(l.ctx, "cache:hot:temp", (in.Page-1)*in.PageSize, in.Page*in.PageSize).Result()
	if err != nil {
		return nil, err
	}

	// 排序
	intSlice, _ := util.ConvertStringsToInts(sorted)
	util.SortByReferenceID(items, intSlice, func(item model.Anime) int {
		return int(item.AnimeId)
	})

	// 删除 cache:hot:temp
	err = l.svcCtx.RedisClient.Del(l.ctx, "cache:hot:temp").Err()

	for _, item := range items {
		resp.AnimeList = append(resp.AnimeList, &pb.Anime{
			AnimeId:     item.AnimeId,
			Title:       item.Title.String,
			Desc:        item.Desc.String,
			Region:      item.Region.String,
			Format:      item.Format.String,
			ImgUrl:      item.ImgUrl.String,
			Studios:     item.Studios.String,
			Status:      item.Status.String,
			Rating:      item.Rating.Float64,
			ReleaseDate: timestamppb.New(item.ReleaseDate.Time),
			UpdateDate:  timestamppb.New(item.UpdateDate.Time),
			UpdatedAt:   timestamppb.New(item.UpdatedAt),
			CreatedAt:   timestamppb.New(item.CreatedAt),
		})
	}
	return resp, nil
}

// 计算热度值的函数，加入时间衰减因素
func CalculateHeatValue(stats model.Stats, currentTime time.Time) float64 {
	// 权重设置（可根据实际情况调整）
	viewWeight := 0.1
	likeWeight := 0.2
	commentWeight := 0.4
	shareWeight := 0.7

	// 设定半衰期（7天）
	halfLife := 7 * 24 * 60 * 60 // 7天的秒数

	// 计算时间差（当前时间 - 创建时间），单位为秒
	timeDifference := currentTime.Sub(stats.UpdatedAt).Seconds()

	// 限制最大时间差为一年的秒数，避免衰减因子过小
	maxTimeDifference := 365 * 24 * 60 * 60 // 1年的秒数
	if timeDifference > float64(maxTimeDifference) {
		timeDifference = float64(maxTimeDifference)
	}

	// 计算时间衰减因子
	decayFactor := math.Exp(-timeDifference / float64(halfLife))

	// 根据权重和衰减因子计算热度值
	heatValue := decayFactor * (float64(stats.ViewCount)*viewWeight +
		float64(stats.LikeCount)*likeWeight +
		float64(stats.CommentCount)*commentWeight +
		float64(stats.ShareCount)*shareWeight)
	return heatValue
}
