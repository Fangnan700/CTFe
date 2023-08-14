package redis

import (
	"CTFe/server/global/config"
	"CTFe/server/util/log"
	"fmt"
	"github.com/go-redis/redis"
)

var (
	rdb *redis.Client
)

func init() {
	var err error

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GlobalConfig.RedisConfig.Host, config.GlobalConfig.RedisConfig.Port),
		Password: config.GlobalConfig.RedisConfig.Password,
		DB:       config.GlobalConfig.RedisConfig.Database,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		log.ErrorLogger.Println(err)
	}
}

func Set(key string, value interface{}) {
	err := rdb.Set(key, value, 0)
	if err != nil {
		log.ErrorLogger.Println(err.Err())
	}
}

func Get(key string) interface{} {
	result, _ := rdb.Get(key).Result()
	return result
}

func Remove(key string) {
	rdb.Del(key)
}
