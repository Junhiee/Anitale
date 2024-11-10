package anime

import (
	"Anitale/apps/anime/rpc/pb"
	"Anitale/pkg/errx"
	"context"
	"github.com/zeromicro/x/errors"

	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnimeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取动画列表分页，根据指定条件进行筛选和排序
func NewGetAnimeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnimeListLogic {
	return &GetAnimeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAnimeListLogic) GetAnimeList(req *types.AnimeListReq) (resp *types.AnimeListResp, err error) {
	resp = &types.AnimeListResp{}
	items, err := l.svcCtx.AnimeRpc.AnimeList(l.ctx, &pb.AnimeListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Tag:      req.Tag,
		Format:   req.Format,
		Region:   req.Region,
		Year:     req.Year,
		Season:   req.Season,
		Sort:     req.Sort,
	})
	if err != nil {
		return nil, errors.New(int(errx.PARAM_ERROR), errx.GetMessage(errx.PARAM_ERROR))
	}

	for _, item := range items.AnimeList {
		tags, err := l.svcCtx.AnimeRpc.GetTags(l.ctx, &pb.TagReq{AnimeId: item.AnimeId})
		if err != nil {
			return nil, errors.New(int(errx.PARAM_ERROR), errx.GetMessage(errx.PARAM_ERROR))
		}
		resp.AnimeList = append(resp.AnimeList, &types.Anime{
			AnimeID:     item.AnimeId,
			Title:       item.Title,
			Desc:        item.Desc,
			ImgURL:      item.ImgUrl,
			Tags:        tags.GetTags(),
			Format:      item.Format,
			Region:      item.Region,
			Rating:      item.Rating,
			Studios:     item.Studios,
			Status:      item.Status,
			ReleaseDate: item.ReleaseDate.String(),
			UpdateDate:  item.UpdateDate.String(),
		})
	}

	resp.Page = items.Page
	resp.PageSize = items.PageSize
	resp.TotalCount = items.TotalCount
	resp.TotalPages = items.TotalPages
	return
}
