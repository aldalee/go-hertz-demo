package client

import (
	"github.com/go-redis/redis/v8"
	"hertz/demo/common/util"
)

type RedisConfig struct {
	GparkModel redis.Options `json:"gpark_model"`
}

type RedisClients struct {
	GparkModel *redis.Client
}

func NewRedisClients(env, configDir string) *RedisClients {
	config := &RedisConfig{}
	if err := util.Load(config, "redis", env, configDir); err != nil {
		panic(err)
	}
	return &RedisClients{
		GparkModel: redis.NewClient(&redis.Options{
			Addr:         config.GparkModel.Addr,
			Password:     config.GparkModel.Password,
			DB:           config.GparkModel.DB,
			PoolSize:     config.GparkModel.PoolSize,
			DialTimeout:  config.GparkModel.DialTimeout,
			ReadTimeout:  config.GparkModel.ReadTimeout,
			WriteTimeout: config.GparkModel.WriteTimeout,
			MinIdleConns: config.GparkModel.MinIdleConns,
		}),
	}
}
