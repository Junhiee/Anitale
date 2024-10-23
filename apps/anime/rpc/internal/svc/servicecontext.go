package svc

import (
	"log"

	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/redis/go-redis/v9"
	zredis "github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"

	"Anitale/apps/anime/rpc/internal/config"
	"Anitale/apps/anime/rpc/model"
)

type ServiceContext struct {
	Config         config.Config
	AnimeModel     model.AnimeModel
	AnimeTagsModel model.AnimeTagsModel
	TagsModel      model.TagsModel
	StatsModel     model.StatsModel
	Conn           *gorm.DB
	CacheClient    *zredis.Redis
	RedisClient    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, err := mysql.Connect(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:         c,
		AnimeModel:     model.NewAnimeModel(conn, c.CacheConf),
		AnimeTagsModel: model.NewAnimeTagsModel(conn, c.CacheConf),
		TagsModel:      model.NewTagsModel(conn, c.CacheConf),
		StatsModel:     model.NewStatsModel(conn, c.CacheConf),
		Conn:           conn,
		CacheClient:    zredis.MustNewRedis(c.RedisConf),
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.RedisConf.Host,
			Password: c.RedisConf.Pass,
			DB:       0}),
	}
}
