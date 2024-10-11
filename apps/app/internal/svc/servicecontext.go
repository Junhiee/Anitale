package svc

import (
	"Anitale/apps/anime/rpc/anime"
	"Anitale/apps/app/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	AnimeRpc anime.Anime
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AnimeRpc: anime.NewAnime(
			zrpc.MustNewClient(c.AnimeRpcConf)),
	}
}
