package global

import (
	"context"
	"github.com/siddontang/go/log"
	"hertz/demo/common/client"
)

var (
	RedisClients *client.RedisClients
)

func Init(env, configDir string) {
	RedisClients = client.NewRedisClients(env, configDir)
	if RedisClients.GparkModel.Ping(context.Background()).Err() != nil {
		log.Errorf("can not connect to redis")
	}
}
