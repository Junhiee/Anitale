package anime

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"Anitale/apps/app/internal/logic/anime"
	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取动画流行趋势数据，根据多个维度（如时间范围、地区、种类、格式
func TrendingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TrendingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := anime.NewTrendingLogic(r.Context(), svcCtx)
		resp, err := l.Trending(&req)
		if err != nil {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
