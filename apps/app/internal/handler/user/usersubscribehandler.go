package user

import (
	"net/http"

	"Anitale/apps/app/internal/logic/user"
	"Anitale/apps/app/internal/svc"
	"Anitale/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户订阅接口
func UserSubscribeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserSubscribeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserSubscribeLogic(r.Context(), svcCtx)
		resp, err := l.UserSubscribe(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
