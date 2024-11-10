package logic

import (
	"context"
	"database/sql"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/model"
	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAnimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAnimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAnimeLogic {
	return &UpdateAnimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新动画信息
func (l *UpdateAnimeLogic) UpdateAnime(in *pb.UpdateAnimeReq) (*pb.UpdateAnimeResp, error) {
	l.svcCtx.AnimeModel.Update(l.ctx, nil, &model.Anime{
		AnimeId: in.AnimeId,
		Title: sql.NullString{
			String: in.Title,
			Valid:  true,
		},
		Desc: sql.NullString{
			String: in.Desc,
			Valid:  true,
		},
		Region: sql.NullString{
			String: in.Region,
			Valid:  true,
		},
		Format: sql.NullString{
			String: in.Format,
			Valid:  true,
		},
		ImgUrl: sql.NullString{
			String: in.ImgUrl,
			Valid:  true,
		},
		Studios: sql.NullString{
			String: in.Studios,
			Valid:  true,
		},
		UpdateDate: sql.NullTime{
			Time:  in.UpdateDate.AsTime(),
			Valid: true,
		},
		ReleaseDate: sql.NullTime{
			Time:  in.ReleaseDate.AsTime(),
			Valid: true,
		},
	})

	return &pb.UpdateAnimeResp{}, nil
}
