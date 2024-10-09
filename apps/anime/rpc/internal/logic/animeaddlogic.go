package logic

import (
	"context"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *AnimeAddLogic) AnimeAdd(in *pb.AnimeAddReq) (*pb.AnimeAddResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AnimeAddResp{}, nil
}
