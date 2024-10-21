package logic

import (
	"context"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAnimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAnimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAnimeLogic {
	return &DeleteAnimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAnimeLogic) DeleteAnime(in *pb.DeleteAnimeReq) (*pb.DeleteAnimeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteAnimeResp{}, nil
}
