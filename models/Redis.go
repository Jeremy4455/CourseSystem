package models

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func RegisterRedis() {
	redisaddr, _ := beego.AppConfig.String("redis::addr")
	redispassword, _ := beego.AppConfig.String("redis::password")
	// 创建 Redis 客户端连接
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     redisaddr,
			Password: redispassword, // 设置密码
			DB:       0,             // 使用默认数据库
		},
	)
}

func RdbGetOrAdd(key string) ([]string, error) {
	exist, err := rdb.Exists(key).Result()
	if err != nil {
		return nil, err
	}
	var values []string
	if exist == 1 {
		values, err := rdb.LRange(key, 0, -1).Result()
		if err != nil {
			return nil, err
		}
		return values, nil
	}
	return values, nil
}
