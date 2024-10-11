package svc

import (
	"Anitale/apps/anime/rpc/animeclient"
	"Anitale/apps/app/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	AnimeRpc animeclient.Anime
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AnimeRpc: animeclient.NewAnime(
			zrpc.MustNewClient(c.AnimeRpcConf)),
	}
}
