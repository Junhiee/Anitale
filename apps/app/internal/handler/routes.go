// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package handler

import (
	"net/http"

	anime "Anitale/apps/app/internal/handler/anime"
	user "Anitale/apps/app/internal/handler/user"
	"Anitale/apps/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/anime",
				Handler: anime.GetAnimeListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
