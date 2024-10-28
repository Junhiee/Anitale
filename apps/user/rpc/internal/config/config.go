package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql   mysql.Mysql
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	RedisConf redis.RedisConf
}
