/*
@Time : 2021/1/7 16:36
@Author : lai
@Description :
@File : main
*/
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
)

//docker build -t go-app-img .
//docker run -d -p 3333:3000 --rm --name go-app-container go-app-img
var ctx = context.Background()

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func test(w http.ResponseWriter, r *http.Request) {
	client := getConnectCluster()
	pong, err := client.Ping(ctx).Result()
	//s, e := client.Get(ctx, "name").Result()
	w.Write([]byte(pong + "\n" + err.Error()))
}

//连接redis
func getConnect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.3.130:6379",
		//Addr: "10.107.207.183:6379",
		//Password: "123456",
		//DB:       0,
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("err:", err)
		fmt.Println(pong)
	}
	fmt.Println("redis connect success")
	return rdb
}

//连接redis集群
func getConnectCluster() *redis.ClusterClient {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"192.168.3.130:6379"}, //set redis cluster url
		Password: "",                             //set password
	})

	return client
}

func main() {
	//getConnect()

	//rdb := getConnect()
	//s, err := rdb.Get(ctx, "name").Result()
	//fmt.Println(err)
	//defer rdb.Close()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(s)
	http.HandleFunc("/", index)
	http.HandleFunc("/test", test)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}
