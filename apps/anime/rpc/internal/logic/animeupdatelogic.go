package logic

import (
	"context"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnimeUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnimeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnimeUpdateLogic {
	return &AnimeUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnimeUpdateLogic) AnimeUpdate(in *pb.AnimeUpdateReq) (*pb.AnimeUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AnimeUpdateResp{}, nil
}
