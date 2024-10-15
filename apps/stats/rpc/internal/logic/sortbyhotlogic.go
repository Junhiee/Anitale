package logic

import (
	"context"

	"Anitale/apps/stats/rpc/internal/svc"
	"Anitale/apps/stats/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SortByHotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSortByHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SortByHotLogic {
	return &SortByHotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 按热度排序, 返回排序完成的 anime_id 列表
func (l *SortByHotLogic) SortByHot(in *pb.SortByHotReq) (*pb.SortByHotResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SortByHotResp{}, nil
}
