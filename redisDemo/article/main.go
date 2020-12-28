/*
@Time : 2020/12/28 16:53
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"
)

const (
	//一星期的秒数
	ONE_WEEK_IN_SCEONDS = 7 * 24 * 60 * 60
	//每张票加的分数
	VOTE_SCORE = 24 * 60 * 60 / 200
)

//投票功能
func articleVote(conn redis.Conn, user, article string) {
	//计算文章投票截止时间
	cutoff := time.Now().Unix() - ONE_WEEK_IN_SCEONDS
	r, err := redis.Int64(conn.Do("ZSCORE", "time:", article))
	if err != nil {
		fmt.Println("ZSCORE time failed,", err)
		return
	}
	//检查是否还可以进行投票，超过7天用户不能再对文章投票
	if cutoff > r {
		return
	}
	//从article:id标识符里面取出文章的id
	article_id := strings.Split(article, ":")[1]

	//如果SADD的命令执行成功，说明用户第一次对文章进行投票，则执行加分操作
	_, err = conn.Do("SADD", "voted"+article_id, user)
	if err != nil {
		fmt.Printf("SADD voted%s err:%v \n", article_id, err)
		return
	}

	//对文章的评分进行更新
	_, err = conn.Do("ZINCRBY", "score:", VOTE_SCORE, article)
	if err != nil {
		fmt.Printf("ZINCRBY score: err:%v \n", err)
		return
	}

	//对文章的投票数量进行更新
	_, err = conn.Do("HINCRBY", "votes", article, 1)
	if err != nil {
		fmt.Printf("ZINCRBY score: err:%v \n", err)
		return
	}

}

//连接redis
func getConnect() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("conn redis failed:", err)
		return nil, err
	}

	fmt.Println("redis conn success")

	defer conn.Close()

	return conn, err
}
func main() {

}
