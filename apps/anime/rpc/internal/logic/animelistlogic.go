package logic

import (
	"context"
	"time"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/model"
	"Anitale/apps/anime/rpc/pb"
	"Anitale/pkg/errx"
	"Anitale/pkg/util"
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

// FIXME 会有全表查询的情况，待优化
func (l *AnimeListLogic) AnimeList(in *pb.AnimeListReq) (*pb.AnimeListResp, error) {
	// 默认值
	if in.Page <= 0 || in.PageSize <= 0 {
		in.Page = 1
		in.PageSize = 10
	}

	var err error
	var resp = &pb.AnimeListResp{}

	var start_date time.Time
	var end_date time.Time
	var start_month time.Time
	var end_month time.Time

	// 用户输入年份，返回该年份的范围
	if in.Year != 0 && in.Season == 0 {
		start_date, end_date, err = util.GetYearRange(int(in.Year))
		if err != nil {
			return nil, errors.Wrapf(errx.NewCustomError(errx.PARAM_ERROR, errx.GetMessage(errx.PARAM_ERROR)), "Func GetYearRange Err: %v", err)
		}
	}

	// 用户输入某个季度 返回改季度的范围
	if in.Year == 0 && in.Season != 0 {
		start_month, end_month, err = util.GetSeasonRange(int(time.Now().Year()), int(in.Season))
		if err != nil {
			return nil, errors.Wrapf(errx.NewCustomError(errx.PARAM_ERROR, errx.GetMessage(errx.PARAM_ERROR)), "Func GetSeasonRange Err: %v", err)
		}
	}

	// 同时有年份和季度, 返回该年份下该季度的范围
	if in.Year != 0 && in.Season != 0 {
		start_date, end_date, err = util.GetSeasonRange(int(in.Year), int(in.Season))
		if err != nil {
			return nil, errors.Wrapf(errx.NewCustomError(errx.PARAM_ERROR, errx.GetMessage(errx.PARAM_ERROR)), "Func GetSeasonRange Err: %v", err)
		}
	}

	var c = model.Condition{
		Region:     in.Region,
		Tag:        in.Tag,
		Format:     in.Format,
		Sort:       in.Sort,
		StartDate:  &start_date,
		EndDate:    &end_date,
		StartMonth: &start_month,
		EndMonth:   &end_month,
	}

	// 查找符合条件的 anime
	items, totalCount, err := l.svcCtx.AnimeModel.PageByCondition(l.ctx, nil, in.Page, in.PageSize, c)
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Filter PageByCond Err: %v", err)
	}

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

	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.TotalCount = totalCount
	resp.TotalPages = (resp.TotalCount + resp.PageSize - 1) / resp.PageSize
	return resp, err
}
