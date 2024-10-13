package svc

import (
	"Anitale/apps/anime/rpc/internal/config"
	"Anitale/apps/anime/rpc/model"
	"log"

	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
)

type ServiceContext struct {
	Config         config.Config
	AnimeModel     model.AnimeModel
	AnimeTagsModel model.AnimeTagsModel
	TagsModel      model.TagsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, err := mysql.Connect(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:     c,
		AnimeModel: model.NewAnimeModel(conn, c.CacheRedis),
		AnimeTagsModel: model.NewAnimeTagsModel(conn, c.CacheRedis),
		TagsModel: model.NewTagsModel(conn, c.CacheRedis),
	}
}
