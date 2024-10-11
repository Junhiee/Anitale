package svc

import (
	"Anitale/apps/app/internal/config"
)

type ServiceContext struct {
	Config config.Config
	// AnimeRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
