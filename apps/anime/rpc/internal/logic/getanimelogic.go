package logic

import (
	"context"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnimeLogic {
	return &GetAnimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnimeLogic) GetAnime(in *pb.GetAnimeReq) (*pb.GetAnimeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetAnimeResp{}, nil
}
