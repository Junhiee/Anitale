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
				// 获取动画角色信息
				Method:  http.MethodGet,
				Path:    "/:anime_id",
				Handler: anime.GetCharacterListHandler(serverCtx),
			},
			{
				// 获取动画剧集信息
				Method:  http.MethodGet,
				Path:    "/episode/:anime_id",
				Handler: anime.GetEpisodeListHandler(serverCtx),
			},
			{
				// 获取动画列表分页，根据指定条件进行筛选和排序
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: anime.GetAnimeListHandler(serverCtx),
			},
			{
				// 获取动画流行趋势数据，根据多个维度（如时间范围、地区、种类、格式
				Method:  http.MethodGet,
				Path:    "/trending",
				Handler: anime.TrendingHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/anime"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 登陆
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginUserHandler(serverCtx),
			},
			{
				// 注册
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 获取个人信息
				Method:  http.MethodGet,
				Path:    "/user/profile/:user_id",
				Handler: user.GetUserProfileHandler(serverCtx),
			},
			{
				// 用户订阅接口
				Method:  http.MethodPost,
				Path:    "/user/subscribe",
				Handler: user.UserSubscribeHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1"),
	)
}
