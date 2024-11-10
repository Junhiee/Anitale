package anime

import (
	"context"

	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCharacterListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取动画角色信息
func NewGetCharacterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCharacterListLogic {
	return &GetCharacterListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCharacterListLogic) GetCharacterList(req *types.GetCharacterReq) (resp *types.GetCharacterListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
