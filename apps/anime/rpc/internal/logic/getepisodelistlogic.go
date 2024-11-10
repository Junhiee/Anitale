package logic

import (
	"context"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEpisodeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEpisodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEpisodeListLogic {
	return &GetEpisodeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Anime Episodes
func (l *GetEpisodeListLogic) GetEpisodeList(in *pb.GetEpisodesReq) (*pb.GetEpisodeListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetEpisodeListResp{}, nil
}
