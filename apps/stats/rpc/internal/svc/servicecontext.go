package svc

import (
	"log"

	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"

	"Anitale/apps/stats/rpc/internal/config"
	"Anitale/apps/stats/rpc/model"
)

type ServiceContext struct {
	Config     config.Config
	StatsModel model.StatsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, err := mysql.Connect(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:     c,
		StatsModel: model.NewStatsModel(conn, c.CacheRedis),
	}
}
