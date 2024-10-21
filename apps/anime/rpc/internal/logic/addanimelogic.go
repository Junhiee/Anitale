package logic

import (
	"context"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAnimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAnimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAnimeLogic {
	return &AddAnimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAnimeLogic) AddAnime(in *pb.AddAnimeReq) (*pb.AddAnimeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddAnimeResp{}, nil
}
