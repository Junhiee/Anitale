package anime

import (
	"net/http"

	"Anitale/apps/app/internal/logic/anime"
	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAnimeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AnimeListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := anime.NewGetAnimeListLogic(r.Context(), svcCtx)
		resp, err := l.GetAnimeList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
