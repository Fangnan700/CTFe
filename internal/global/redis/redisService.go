package redis

import (
	"CTFe/internal/global/config"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
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
		fmt.Println(err)
	}
}

// Remove 删除redis key
func Remove(key string) {
	rdb.Del(key)
}

// SetCTFeTokenStatus 设置CTFeToken状态
func SetCTFeTokenStatus(ctfeToken string, status int) {
	err := rdb.Set(ctfeToken, status, 0)
	fmt.Println(err)
}

// GetCTFeTokenStatus 获取CTFeToken状态
func GetCTFeTokenStatus(ctfeToken string) int {
	result, _ := rdb.Get(ctfeToken).Result()
	status, _ := strconv.Atoi(result)
	return status
}
