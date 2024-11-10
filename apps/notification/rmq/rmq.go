package main

import (
	"flag"
	"fmt"

	"Anitale/apps/notification/rmq/internal/config"
	"Anitale/apps/notification/rmq/internal/listen"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

// TODO asynq server

var configFile = flag.String("f", "etc/rmq.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()

	fmt.Printf("Starting mq server at %s...\n", c.ListenOn)

}
