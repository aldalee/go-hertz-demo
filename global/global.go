package global

import (
	"context"
	"flag"
	"github.com/siddontang/go/log"
	"hertz/demo/common/client"
)

var (
	RedisClients *client.RedisClients
)

func Init() {
	var env, configDir string
	flag.StringVar(&env, "env", "dev", "config env name")
	flag.StringVar(&configDir, "config_dir", "./common/config", "config file dir")
	flag.Parse()
	log.Infof("gpark_mms starting... env: %s, config_dir: %s", env, configDir)

	RedisClients = client.NewRedisClients(env, configDir)
	if RedisClients.GparkModel.Ping(context.Background()).Err() != nil {
		log.Errorf("can not connect to redis")
	}
}
