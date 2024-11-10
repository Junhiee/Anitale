package svc

import (
	"Anitale/apps/anime/rpc/animeservice"
	"Anitale/apps/app/internal/config"
	"Anitale/apps/user/rpc/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	AnimeRpc animeservice.AnimeService
	UserRpc  userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		AnimeRpc: animeservice.NewAnimeService(zrpc.MustNewClient(c.AnimeRpcConf)),
		UserRpc:  userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
