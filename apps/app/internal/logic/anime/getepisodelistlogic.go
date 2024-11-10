package anime

import (
	"context"

	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEpisodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取动画剧集信息
func NewGetEpisodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEpisodeListLogic {
	return &GetEpisodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEpisodeListLogic) GetEpisodeList(req *types.GetEpisodeReq) (resp *types.GetEpisodeListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
