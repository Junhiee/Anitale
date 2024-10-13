package logic

import (
	"context"
	"fmt"
	"time"

	"Anitale/apps/anime/rpc/internal/svc"
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
	// 默认值
	if in.Page <= 0 || in.PageSize <= 0 {
		in.Page = 1
		in.PageSize = 20
	}
	if in.Year == 0 {
		in.Year = int64(time.Now().Year())
	}

	var resp = &pb.AnimeListResp{}

	var anime_ids []int64
	if in.Tag != "" {
		// 查找 tag_id
		tag_id, err := l.svcCtx.TagsModel.FindOneByTag(l.ctx, in.Tag)
		if err != nil {
			return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Filter FindOneByTag Err: %v", err)
		}

		// 查找符合 tag_id 的 anime_id
		anime_tags, err := l.svcCtx.AnimeTagsModel.PageByTags(l.ctx, tag_id, nil)
		if err != nil {
			return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Filter PageByTags Err: %v", err)
		}

		for _, anime_tag := range anime_tags {
			anime_ids = append(anime_ids, anime_tag.AnimeId)
		}
	}

	// 查找符合季度的 anime
	// 用户输入 年份 返回该年份的所有季度的 anime
	start_date, end_date, err := GetYearRange(int(in.Year))
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomError(errx.PARAM_ERROR, errx.GetMessage(errx.PARAM_ERROR)), "Func GetYearRange Err: %v", err)
	}

	// TODO 用户输入 季度 返回改季度的所有 anime

	// 同时有年份和季度, 返回该年份下该季度的 anime
	if in.Season != 0 {
		start_date, end_date, err = GetSeasonRange(int(in.Year), int(in.Season))
		if err != nil {
			return nil, errors.Wrapf(errx.NewCustomError(errx.PARAM_ERROR, errx.GetMessage(errx.PARAM_ERROR)), "Func GetSeasonRange Err: %v", err)
		}
	}

	// 查询条件
	conditon := map[string]interface{}{
		"region":     in.Region,
		"anime_type": in.AnimeType,
		"start_date": start_date,
		"end_date":   end_date,
	}

	// 查找符合条件的 anime
	items, err := l.svcCtx.AnimeModel.PageByCond(l.ctx, conditon, in.Page, in.PageSize, anime_ids, nil)
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Filter PageByCond Err: %v", err)
	}

	for _, item := range items {
		resp.AnimeList = append(resp.AnimeList, &pb.Item{
			AnimeId:     item.AnimeId,
			Title:       item.Title.String,
			Desc:        item.Desc.String,
			Region:      item.Region.String,
			AnimeType:   item.AnimeType.String,
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
	return resp, err
}

// 输入年份和季度,返回该年份下该季度的范围
func GetSeasonRange(year int, season int) (time.Time, time.Time, error) {
	switch season {
	case 1:
		return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(year, time.March, 31, 23, 59, 59, 0, time.UTC), nil
	case 2:
		return time.Date(year, time.April, 1, 0, 0, 0, 0, time.UTC),
			time.Date(year, time.June, 30, 23, 59, 59, 0, time.UTC), nil
	case 3:
		return time.Date(year, time.July, 1, 0, 0, 0, 0, time.UTC),
			time.Date(year, time.September, 30, 23, 59, 59, 0, time.UTC), nil
	case 4:
		return time.Date(year, time.October, 1, 0, 0, 0, 0, time.UTC),
			time.Date(year, time.December, 31, 23, 59, 59, 0, time.UTC), nil
	default:
		return time.Time{}, time.Time{}, fmt.Errorf("无效的季度: %d", season)
	}
}

func GetYearRange(year int) (time.Time, time.Time, error) {
	nextYear := time.Now().AddDate(1, 0, 0).Year()
	if year < 1900 || year > nextYear {
		return time.Time{}, time.Time{}, fmt.Errorf("年份不合理，必须在 1900 到当前年份的后一年之间: %d", nextYear)
	}

	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(year, time.December, 31, 23, 59, 59, 0, time.UTC), nil

}
