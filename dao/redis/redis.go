package redis

import (
	"bluebell/settings"
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

var ctx = context.Background()

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

func Set(key string, value string) (err error) {
	err = rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		zap.Error(err)
		return err
	}
	return
}

func Get(key string) (value string, err error) {
	value, err = rdb.Get(ctx, key).Result()
	if err != nil {
		return "", errors.New("redis获取失败")
	} else {
		return value, nil
	}
}
