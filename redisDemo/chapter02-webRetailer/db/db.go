/*
@Time : 2020/12/29 16:22
@Author : lai
@Description :
@File : db
*/
package db

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

//连接redis
func GetConnect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.130:6379",
		Password: "123456",
		DB:       0,
	})
	fmt.Println("redis connect success")
	return rdb
}
