package logic

import (
	"Anitale/pkg/errx"
	"context"
	"github.com/pkg/errors"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagsLogic {
	return &GetTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get Anime Tag List
func (l *GetTagsLogic) GetTags(in *pb.TagReq) (*pb.TagResp, error) {
	var resp = &pb.TagResp{}
	items, err := l.svcCtx.AnimeTagsModel.FindTagIdByAnimeId(l.ctx, in.AnimeId)
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "FindTagIdByAnimeId Err: %v", err)
	}

	for _, item := range items {
		tag, err := l.svcCtx.TagsModel.FindOne(l.ctx, item.TagId)
		if err != nil {
			return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "FindOne Err: %v", err)
		}

		resp.Tags = append(resp.Tags, tag.Tag)
	}

	return resp, nil
}
