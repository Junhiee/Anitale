package logic

import (
	"context"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/model"
	"Anitale/apps/anime/rpc/pb"
	"Anitale/pkg/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type AnimeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnimeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnimeListLogic {
	return &AnimeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnimeListLogic) AnimeList(in *pb.AnimeListReq) (*pb.AnimeListResp, error) {
	// 默认页码和默认每页大小
	if in.PageSize <= 0 {
		in.PageSize = 2
	}

	if in.Page <= 0 {
		in.Page = 1
	}

	var err error
	var resp = &pb.AnimeListResp{}
	var items = []model.Anime{}

	// 默认分页, 通过 update_at 排序
	items, err = l.svcCtx.AnimeModel.AnimeList(l.ctx, in.Page, in.PageSize, nil)
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Query GetPageAnimeList Err: %v", err)
	}

	// 通过 tag 分页, 按 update_at 排序
	if in.Tag != "" {
		items, err = l.svcCtx.AnimeModel.PageByTag(l.ctx, in.Tag, in.Page, in.PageSize, nil)
		if err != nil {
			return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Filter PageByTag Err: %v", err)
		}
	}

	// 通过 country 分页, 按 update_at 排序
	if in.Country != "" {
		items, err = l.svcCtx.AnimeModel.PageByCountry(l.ctx, in.Country, in.Page, in.PageSize, nil)
		if err != nil {
			return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Filter PageByCountry Err: %v", err)
		}
	}

	// TODO 按季度分页 - 4月 7月 10月 12月
	// 通过 relase_date 分页, 按 update_at 排序
	if in.RelaseDate != nil {
		items, err = l.svcCtx.AnimeModel.PageByRelaseDate(l.ctx, in.RelaseDate.AsTime(), in.Page, in.PageSize, nil)
		if err != nil {
			return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Filter PageByRelaseDate Err: %v", err)
		}
	}

	// 通过 anime_type 分页, 按 update_at 排序
	if in.AnimeType != "" {
		items, err = l.svcCtx.AnimeModel.PageByAnimeType(l.ctx, in.AnimeType, in.Page, in.PageSize, nil)
		if err != nil {
			return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Filter PageByAnimeType Err: %v", err)
		}
	}

	for _, item := range items {
		resp.AnimeList = append(resp.AnimeList, &pb.Item{
			AnimeId:    item.AnimeId,
			Title:      item.Title.String,
			Desc:       item.Desc.String,
			Country:    item.Country.String,
			AnimeType:  item.AnimeType.String,
			Tag:        item.Tag.String,
			Studios:    item.Studios.String,
			Status:     item.Status.String,
			Rating:     item.Rating.Int64,
			RelaseDate: timestamppb.New(item.RelaseDate.Time),
		})
	}
	return resp, err
}
