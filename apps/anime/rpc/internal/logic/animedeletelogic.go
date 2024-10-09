package logic

import (
	"context"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnimeDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnimeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnimeDeleteLogic {
	return &AnimeDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnimeDeleteLogic) AnimeDelete(in *pb.AnimeDeleteReq) (*pb.AnimeDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AnimeDeleteResp{}, nil
}
