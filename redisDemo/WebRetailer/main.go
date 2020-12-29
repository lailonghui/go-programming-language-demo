/*
@Time : 2020/12/29 16:20
@Author : lai
@Description :
@File : main
*/
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"math"
	"strings"
	"time"
)

const (
	QUIT = false
	//保存的会话数量
	LIMIT = 10000
)

var ctx = context.Background()

//检查用户是否登录,如果登录则返回用户id
func checkToken(rdb *redis.Client, token string) string {
	return rdb.HGet(ctx, "login:", token).Val()
}

//更新令牌
func updateToken(rdb *redis.Client, token, user, item string) {
	timestamp := float64(time.Now().Unix())

	//维持令牌和用户的映射
	rdb.HSet(ctx, "login:", token, user)

	//记录令牌最后一次出现的时间
	rdb.ZAdd(ctx, "recent:", &redis.Z{Score: timestamp, Member: token})

	if len(strings.TrimSpace(item)) != 0 {
		//记录用户浏览过的商品,item为商品id
		rdb.ZAdd(ctx, "viewed:"+token, &redis.Z{Score: timestamp, Member: item})
		//移除旧的记录,只保留用户最近浏览过的25个商品
		rdb.ZRemRangeByRank(ctx, "viewed:"+token, 0, -26)
	}
}

//清理会话数据
func clearSession(rdb *redis.Client) {
	for !QUIT {
		size := rdb.ZCard(ctx, "recent:").Val()
		if size <= LIMIT {
			time.Sleep(time.Second)
			continue
		}

		endIndex := math.Min(float64(size-LIMIT), 100)
		//获取需要移除的令牌id
		tokens := rdb.ZRange(ctx, "recent:", 0, int64(endIndex)).Val()

		var sessionKeys []string
		for _, token := range tokens {
			sessionKeys = append(sessionKeys, "viewed:"+token)
		}

		//移除最旧的令牌
		rdb.Del(ctx, sessionKeys...)
		rdb.HDel(ctx, "login:", sessionKeys...)
		rdb.ZRem(ctx, "recent:", sessionKeys...)
	}

}

func main() {

}
