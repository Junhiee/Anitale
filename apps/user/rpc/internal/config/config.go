package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Mysql mysql.Mysql
	zrpc.RpcServerConf
}
