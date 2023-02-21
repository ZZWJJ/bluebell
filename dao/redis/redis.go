package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Init(redisConfig *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password, // 密码
		DB:       redisConfig.Db,       // 数据库
		PoolSize: redisConfig.PoolSize, // 连接池大小
	})

	_, err = rdb.Ping(rdb.Context()).Result()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Close() {
	_ = rdb.Close()
}
