package svc

import (
	"Anitale/apps/anime/rpc/anime"
	"Anitale/apps/app/internal/config"
	"Anitale/apps/user/rpc/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	AnimeRpc anime.Anime
	UserRpc  userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		AnimeRpc: anime.NewAnime(zrpc.MustNewClient(c.AnimeRpcConf)),
		UserRpc:  userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
