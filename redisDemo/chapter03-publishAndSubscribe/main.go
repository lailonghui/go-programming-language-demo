/*
@Time : 2020/12/30 15:22
@Author : lai
@Description :
@File : main
*/
package main

import (
	"context"
	"fmt"
	"lai.com/go_programming_language_demo/redisDemo/chapter03-publishAndSubscribe/db"
)

var ctx = context.Background()

func main() {
	rdb := db.GetConnect()
	data, err := rdb.Get(ctx, "test").Result()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(111)
	fmt.Println(data)
	fmt.Println(222)
	//info := rdb.Info(ctx, "aof_pending_bio_fsync").Val()
	//fmt.Println(info)
	//rdb.Publish(ctx, "c1", "hello! what's your name?")
}
