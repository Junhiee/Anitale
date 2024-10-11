package anime

import (
	"context"
	"fmt"

	"Anitale/apps/anime/rpc/pb"
	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnimeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAnimeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnimeListLogic {
	return &GetAnimeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAnimeListLogic) GetAnimeList(req *types.AnimeListReq) (resp *types.AnimeListResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.AnimeRpc.AnimeList(l.ctx, &pb.AnimeListReq{
		Page:      req.Page,
		PageSize:  int64(req.PageSize),
		AnimeType: req.AnimeType,
	})

	fmt.Println(res)

	resp = &types.AnimeListResp{}

	resp.Code = 0
	resp.Msg = "success"
	items := res.GetAnimeList()
	fmt.Println(items)
	// for _, item := range items {
	// 	resp.Data = append(resp.Data, &types.Anime{
	// 		Title:      item.Title,
	// 		ImgURL:     "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2527307692.jpg",
	// 		Tag:        item.Tag,
	// 		UpdateTime: item.UpdatedAt.Seconds,
	// 	})
	// }
	resp.Data = nil
	return resp, err
}
