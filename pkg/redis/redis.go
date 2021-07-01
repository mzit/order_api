package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/mzit/order_api/pkg/setting"
	"time"
)

// 声明一个全局的rdb变量
var rdb *redis.Client
var ctx = context.Background()

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
		DB:       0, // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func Set(key string, data interface{}, second int) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = rdb.Set(ctx, key, value, time.Duration(second)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
