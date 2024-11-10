package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/model"
	"Anitale/apps/anime/rpc/pb"
	"Anitale/pkg/errx"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEpisodesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEpisodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEpisodesLogic {
	return &UpdateEpisodesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新动画剧集信息
// 当动画剧集更新时，将消息推送至消息队列
func (l *UpdateEpisodesLogic) UpdateEpisodes(in *pb.UpdateEpisodesReq) (*pb.UpdateEpisodesResp, error) {
	// 动画剧集episodes表更新
	err := l.svcCtx.EpisodesModel.Update(l.ctx, nil, &model.Episodes{
		EpisodeId: in.EpisodeId,
		AnimeId: sql.NullInt64{
			Int64: in.AnimeId,
			Valid: true,
		},
		EpisodeNumber: sql.NullInt32{
			Int32: in.EpisodeNumber,
			Valid: true,
		},
		Title: sql.NullString{
			String: in.Title,
			Valid:  true,
		},
		ReleaseDate: sql.NullTime{
			Time:  in.ReleaseDate.AsTime(),
			Valid: true,
		},
		Duration: sql.NullInt32{
			Int32: in.Duration,
			Valid: true,
		},
		Synopsis: sql.NullString{
			String: in.Synopsis,
			Valid:  true,
		},
		VideoUrl: sql.NullString{
			String: in.VideoUrl,
			Valid:  true,
		},
	})
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Update episodes Err: %v", err)
	}

	// 更新anime表中动画的更新日期字段
	err = l.svcCtx.AnimeModel.Update(l.ctx, nil, &model.Anime{
		AnimeId: in.AnimeId,
		UpdateDate: sql.NullTime{
			Time:  in.ReleaseDate.AsTime(),
			Valid: true,
		},
	})
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Update anime Err: %v", err)
	}

	// 将更新消息事件保存至 anime_updates 表
	var animeUpdates = model.AnimeUpdates{
		AnimeId: in.AnimeId,
		EpisodeId: sql.NullInt64{
			Int64: in.EpisodeId,
			Valid: true,
		},
		UpdateType: "new_episode",
		UpdateDescription: sql.NullString{
			String: "新剧集发布",
			Valid:  true,
		},
	}
	err = l.svcCtx.AnimeUpdatesModel.Insert(l.ctx, nil, &animeUpdates)
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Insert anime_updates Err: %v", err)
	}

	// 将该更新事件推送至消息队列
	task, err := WebEpisodesTask(&animeUpdates)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.AsynqClient.EnqueueContext(l.ctx, task)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateEpisodesResp{}, nil
}

const (
	TypeEmailDelivery = "web:episodes"
)

func WebEpisodesTask(data *model.AnimeUpdates) (*asynq.Task, error) {
	p, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, p), nil
}
