package logic

import (
	"context"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddEpisodesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddEpisodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddEpisodesLogic {
	return &AddEpisodesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddEpisodesLogic) AddEpisodes(in *pb.AddEpisodesReq) (*pb.AddEpisodesResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddEpisodesResp{}, nil
}
