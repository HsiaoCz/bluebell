package redis

import (
	"bluebell/setting"
	"fmt"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init() (err error) {
	rcg := setting.Conf.RedisConfig
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rcg.Host, rcg.Port),
		Password: rcg.Password,
		DB:       rcg.DB,
		PoolSize: rcg.PoolSize,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		zap.L().Error("redis ping err:", zap.Error(err))
		return
	}
	return err
}

func Close() {
	_ = rdb.Close()
}
