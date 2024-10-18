package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"
)

type AnimeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnimeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnimeAddLogic {
	return &AnimeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnimeAddLogic) AnimeAdd(in *pb.AddAnimeReq) (*pb.AddAnimeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddAnimeResp{}, nil
}
