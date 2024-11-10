package svc

import (
	"log"

	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	zredis "github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"

	"Anitale/apps/anime/rpc/internal/config"
	"Anitale/apps/anime/rpc/model"
)

type ServiceContext struct {
	Config            config.Config
	AnimeModel        model.AnimeModel
	AnimeTagsModel    model.AnimeTagsModel
	TagsModel         model.TagsModel
	StatsModel        model.StatsModel
	EpisodesModel     model.EpisodesModel
	AnimeUpdatesModel model.AnimeUpdatesModel

	Conn        *gorm.DB
	CacheClient *zredis.Redis
	RedisClient *redis.Client
	AsynqClient *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, err := mysql.Connect(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceContext{
		Config:            c,
		AnimeModel:        model.NewAnimeModel(conn, c.CacheConf),
		AnimeTagsModel:    model.NewAnimeTagsModel(conn, c.CacheConf),
		TagsModel:         model.NewTagsModel(conn),
		StatsModel:        model.NewStatsModel(conn, c.CacheConf),
		EpisodesModel:     model.NewEpisodesModel(conn, c.CacheConf),
		AnimeUpdatesModel: model.NewAnimeUpdatesModel(conn, c.CacheConf),

		Conn:        conn,
		CacheClient: zredis.MustNewRedis(c.RedisConf),
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.RedisConf.Host,
			Password: c.RedisConf.Pass,
			DB:       0}),

		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}),
	}
}
