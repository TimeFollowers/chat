package initialize

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"wutool.cn/chat/server/global"
)

func InitRedis() *redis.Client {
	r := global.CONFIG.Redis
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     r.Addr,     // 要连接的redis IP:port
		Password: r.Password, // redis 密码
		DB:       r.DB,       // 要连接的redis 库
	})
	// 检测心跳
	pong, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("connect redis failed")
		panic("redis连接失败, err:" + err.Error())
	}
	fmt.Printf("redis ping result: %s\n", pong)
	return RedisClient
}
