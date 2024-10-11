package anime

import (
	"context"

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

	return
}
