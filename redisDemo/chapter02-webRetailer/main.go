/*
@Time : 2020/12/29 16:20
@Author : lai
@Description :
@File : main
*/
package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"lai.com/go_programming_language_demo/redisDemo/chapter02-webRetailer/db"
	"math"
	"net/http"
	"reflect"
	"strings"
	"time"
)

const (
	QUIT = false
	//保存的会话数量
	LIMIT = 10000
)

var ctx = context.Background()

//将传入slice的每个元素拿出来interface()化
func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic(" arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

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
		//记录所有商品被浏览的次数
		rdb.ZIncrBy(ctx, "viewed:", 1, item)
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
			sessionKeys = append(sessionKeys, "cart:"+token)
		}

		//移除最旧的令牌
		rdb.Del(ctx, sessionKeys...)
		rdb.HDel(ctx, "login:", sessionKeys...)
		rdb.ZRem(ctx, "recent:", ToSlice(sessionKeys))
	}

}

//更新购物车数据
func addToCart(rdb *redis.Client, token, item string, count int) {
	if count <= 0 {
		//从购物车移除指定的商品
		rdb.HDel(ctx, "cart:"+token, item)
	} else {
		//将指定的商品添加到购物车
		rdb.HSet(ctx, "cart:"+token, item, count)
	}
}

//缓存页面
func cacheRequest(rdb *redis.Client, req http.Request, callback func(req http.Request) string) string {

	hash := sha256.Sum256([]byte(req.URL.String()))
	pageKey := "cache:" + fmt.Sprintf("%X", hash)
	content := rdb.Get(ctx, pageKey).Val()

	if len(strings.TrimSpace(content)) == 0 {
		content = callback(req)
		rdb.Set(ctx, pageKey, content, 300)
	}
	return content
}

type Student struct {
	Name  string
	Age   int
	Hobby string
}

//调度缓存和终止缓存的函数
func scheduleRowCache(rdb *redis.Client, rowId string, delay float64) {
	//先设置数据行的延迟值
	rdb.ZAdd(ctx, "delay:", &redis.Z{Score: delay, Member: rowId})
	//立即对需要缓存的数据行进行调度
	rdb.ZAdd(ctx, "schedule:", &redis.Z{Score: float64(time.Now().Unix()), Member: rowId})
}

//数据行缓存(守护进程函数)
func cacheRows(rdb *redis.Client) {
	for !QUIT {
		//获取调度有序集合的第一个元素
		next, err := rdb.ZRangeWithScores(ctx, "schedule:", 0, 0).Result()
		now := float64(time.Now().Unix())
		//如果暂时没有行需要被缓存，则休眠50毫秒
		if err != nil || next[0].Score > now {
			time.Sleep(time.Millisecond * 50)
		}

		rowId := next[0].Member.(string)
		//获取下一次调度的延迟时间
		delay := rdb.ZScore(ctx, "delay:", rowId).Val()
		if delay <= 0 {
			rdb.ZRem(ctx, "delay:", rowId)
			rdb.ZRem(ctx, "schedule:", rowId)
			rdb.Del(ctx, "inv:"+rowId)
		}

		//更新调度时间并设置缓存值
		rdb.ZAdd(ctx, "schedule:", &redis.Z{Score: now + delay, Member: rowId})

		s := &Student{
			Name:  "lai",
			Age:   20,
			Hobby: "basketball",
		}
		bytes, _ := json.Marshal(s)
		rdb.Set(ctx, "inv:"+rowId, string(bytes), -1)
	}
}

//重新调整浏览次数函数(守护进程函数)
func rescaleViewed(rdb *redis.Client) {
	for !QUIT {
		rdb.ZRemRangeByRank(ctx, "viewed", 0, -20001)
		rdb.ZInterStore(ctx, "viewed:", &redis.ZStore{
			Keys:    []string{"viewed:"},
			Weights: []float64{.5},
		})
		time.Sleep(time.Second * 300)
	}
}

func main() {

	//var date1 []byte = []byte("adfkfjsadsijfal")
	//var hs = sha256.Sum256(date1)
	//fmt.Printf("%X\n", hs)
	rdb := db.GetConnect()
	//rdb.ZInterStore(ctx, "viewed:", &redis.ZStore{
	//	Keys:    []string{"viewed:"},
	//	Weights: []float64{0.5},
	//})
}
