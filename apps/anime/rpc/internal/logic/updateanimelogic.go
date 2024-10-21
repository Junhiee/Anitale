package logic

import (
	"context"

	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAnimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAnimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAnimeLogic {
	return &UpdateAnimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAnimeLogic) UpdateAnime(in *pb.UpdateAnimeReq) (*pb.UpdateAnimeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateAnimeResp{}, nil
}
