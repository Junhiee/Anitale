package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      mysql.Mysql
	CacheRedis cache.CacheConf
}
