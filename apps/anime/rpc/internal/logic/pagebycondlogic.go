package logic

import (
	"context"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"
	"Anitale/pkg/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PageByCondLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPageByCondLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageByCondLogic {
	return &PageByCondLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 多条件分页查询

func (l *PageByCondLogic) PageByCond(in *pb.PageByCondReq) (*pb.PageByCondResp, error) {

	if in.PageSize <= 0 {
		in.PageSize = 2
	}

	if in.Page <= 0 {
		in.Page = 1
	}

	var resp = &pb.PageByCondResp{}

	cond := map[string]string{
		"country":    in.Country,
		"tag":        in.Tag,
		"anime_type": in.AnimeType,
	}

	items, err := l.svcCtx.AnimeModel.PageByCond(l.ctx, cond, in.Page, in.PageSize, nil)
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomError(errx.DB_ERROR, errx.GetMessage(errx.DB_ERROR)), "Filter PageByCond Err: %v", err)
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
