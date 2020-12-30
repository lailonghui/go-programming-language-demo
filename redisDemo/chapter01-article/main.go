/*
@Time : 2020/12/28 16:53
@Author : lai
@Description :
@File : main
*/
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

var ctx = context.Background()

const (
	//一星期的秒数
	ONE_WEEK_IN_SCEONDS = 7 * 24 * 60 * 60
	//每张票加的分数
	VOTE_SCORE = 24 * 60 * 60 / 200
	//每次查询的页数
	ARTICLE_PER_PAGE = 25
)

//连接redis
func getConnect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.130:6379",
		Password: "123456",
		DB:       0,
	})
	fmt.Println("redis connect success")
	return rdb
}

//投票功能
func articleVote(rdb *redis.Client, user, article string) {
	//计算文章投票截止时间
	cutoff := float64(time.Now().Unix() - ONE_WEEK_IN_SCEONDS)
	r, err := rdb.ZScore(ctx, "time:", article).Result()
	if err != nil {
		fmt.Println("ZSCORE time failed,", err)
		return
	}
	//检查是否还可以进行投票，超过7天用户不能再对文章投票
	if cutoff > r {
		return
	}
	//从article:id标识符里面取出文章的id
	articleId := strings.Split(article, ":")[1]

	//如果SADD的命令执行成功，说明用户第一次对文章进行投票，则执行加分操作
	result, err := rdb.SAdd(ctx, "voted"+articleId, user).Result()
	if result == 0 {
		fmt.Println("you already voted")
		return
	}
	if err != nil {
		fmt.Printf("SADD voted%s err:%v \n", articleId, err)
		return
	}

	//对文章的评分进行更新
	err = rdb.ZIncrBy(ctx, "score:", VOTE_SCORE, article).Err()
	if err != nil {
		fmt.Printf("ZINCRBY score: err:%v \n", err)
		return
	}

	//对文章的投票数量进行更新
	err = rdb.HIncrBy(ctx, article, "votes", 1).Err()
	if err != nil {
		fmt.Printf("HINCRBY votes: err:%v \n", err)
		return
	}

}

//发布文章
func postArticle(rdb *redis.Client, user, title, link string) int64 {
	//使用INCR累加生成article_id
	articleId, err := rdb.Incr(ctx, "chapter01-article:").Result()
	if err != nil {
		fmt.Printf("INCR chapter01-article: err:%s\n", err)
		return 0
	}

	//把文章发布者的id添加到记录文章已投票的集合中
	voted := fmt.Sprintf("voted:%d", articleId)
	err = rdb.SAdd(ctx, voted, user).Err()
	if err != nil {
		fmt.Printf("SADD voted: %s %s err:%v\n", voted, user, err)
		return 0
	}
	//设置过期时间
	err = rdb.Expire(ctx, voted, ONE_WEEK_IN_SCEONDS).Err()
	if err != nil {
		fmt.Printf("EXPIRE %s %d err:%v\n", voted, ONE_WEEK_IN_SCEONDS, err)
		return 0
	}

	//将文章信息存储到一个散列里面
	now := time.Now().Unix()
	article := fmt.Sprintf("chapter01-article:%d", articleId)
	err = rdb.HMSet(ctx, article, "title", title, "link", link, "poster", user, "time", now, "votes", 1).Err()
	if err != nil {
		fmt.Printf("HMSET %s failed,err:%v\n", article, err)
		return 0
	}

	//将文章添加到根据评分排序的有序集合里
	err = rdb.ZAdd(ctx, "score:", &redis.Z{Score: float64(now + VOTE_SCORE), Member: article}).Err()
	if err != nil {
		fmt.Printf("ZADD score: %s failed,err:%v\n", article, err)
		return 0
	}

	//将文章添加到根据发布时间排序的有序集合里
	err = rdb.ZAdd(ctx, "time:", &redis.Z{Score: float64(now), Member: article}).Err()
	if err != nil {
		fmt.Printf("ZADD time: %s failed,err:%v\n", article, err)
		return 0
	}

	return articleId
}

//获取文章
func getArticle(rdb *redis.Client, page int64, order string) []map[string]string {
	//获取文章的起始索引和结束索引
	start := (page - 1) * ARTICLE_PER_PAGE
	end := start + ARTICLE_PER_PAGE - 1

	//获取多个文章的id
	ids, err := rdb.ZRevRange(ctx, order, start, end).Result()
	if err != nil {
		panic(err)
	}

	var articles []map[string]string
	//根据文章id获取文章详细信息
	for _, id := range ids {
		articleData := rdb.HGetAll(ctx, id).Val()
		articleData["id"] = id
		articles = append(articles, articleData)
	}
	return articles
}

//添加和移除群组里的文章
func addRemoveGroup(rdb *redis.Client, articleId int64, toAdd, toRemove []string) {
	article := fmt.Sprintf("chapter01-article:%d", articleId)
	for _, group := range toAdd {
		rdb.SAdd(ctx, "group:"+group, article)
	}
	for _, group := range toRemove {
		rdb.SAdd(ctx, "group:"+group, article)
	}
}

//从群组里获取一整夜文章
func getGroupArticles(rdb *redis.Client, group string, page int64, order string) []map[string]string {
	key := order + group
	fmt.Println(key)
	//检查是否有已缓存的排序结果,如果没有再进行排序
	if rdb.Exists(ctx, key).Val() == 0 {
		//根据评分或发布时间,对群组文章进行排序
		err := rdb.ZInterStore(ctx, key, &redis.ZStore{
			Keys:      []string{"group:" + group, order},
			Aggregate: "max",
		}).Err()
		if err != nil {
			panic(err)
		}
		err = rdb.Expire(ctx, key, time.Second).Err()
		if err != nil {
			panic(err)
		}
	}
	return getArticle(rdb, page, order)
}

func main() {

	rdb := getConnect()

	//articleId := postArticle(rdb, "1", "testTitle02", "http://lai.com")
	//fmt.Println("articleId=", articleId)

	//articles := getArticle(rdb, 1, "score:")
	//for _, v := range articles {
	//	fmt.Println(v)
	//}

	//articleVote(rdb, "2", "chapter01-article:10")

	//addRemoveGroup(rdb, 11, []string{"test"}, nil)

	articles := getGroupArticles(rdb, "test", 1, "score:")
	for _, v := range articles {
		fmt.Println(v)
	}

}
