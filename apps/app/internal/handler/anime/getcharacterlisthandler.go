package anime

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"Anitale/apps/app/internal/logic/anime"
	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取动画角色信息
func GetCharacterListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCharacterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := anime.NewGetCharacterListLogic(r.Context(), svcCtx)
		resp, err := l.GetCharacterList(&req)
		if err != nil {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
