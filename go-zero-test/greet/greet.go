package main

import (
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	"greet/common/middleware"
	"time"

	"greet/internal/config"
	"greet/internal/handler"
	"greet/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	server.Use(middleware.NewAuthMiddleware().Handle)
	server.Use(middleware.NewRequestMiddleware().Handle)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

	ticker := time.NewTicker(1 * time.Second)
	rdb := redis.NewClient(&redis.Options{})
	rdb.ZRangeByScore()
}
