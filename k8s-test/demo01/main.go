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

//docker build -t 192.168.3.130:5000/go-app:v12 .
//docker push 192.168.3.130:5000/go-app:v12

//kubectl create -f go-app.yaml
//kubectl expose deploy go-app --port=3000 --type=NodePort

//kubectl get svc --all-namespaces -o wide

//kubectl delete deploy go-app && kubectl delete svc go-app
var ctx = context.Background()

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World 1209</h1>")
}

func test(w http.ResponseWriter, r *http.Request) {
	client := getConnectCluster()
	name, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Fprintln(w, "err=", err.Error())
	}
	fmt.Fprintln(w, name)
}

func set(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	value := r.FormValue("value")
	fmt.Println(name)
	fmt.Println(value)
	client := getConnectCluster()
	res, err := client.Set(ctx, name, value, -1).Result()
	if err != nil {
		fmt.Fprintln(w, "err=", err.Error())
	}
	fmt.Fprintln(w, res)
}

func get(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	client := getConnectCluster()
	res, err := client.Get(ctx, name).Result()
	if err != nil {
		fmt.Fprintln(w, "err=", err.Error())
	}
	fmt.Fprintln(w, name+"=", res)
}

//连接redis
func getConnect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.3.132:7001",
		//Addr: "192.168.3.130:6379",
		//Addr: "10.97.213.221:6379",
		//Password: "123456",
		//DB:       0,
	})
	//pong, err := rdb.Ping(ctx).Result()
	//if err != nil {
	//	fmt.Println("err:", err)
	//	fmt.Println(pong)
	//}
	fmt.Println("redis connect success")
	return rdb
}

//连接redis集群
func getConnectCluster() *redis.ClusterClient {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		//Addrs: []string{"192.168.3.132:7001", "192.168.3.132:7002", "192.168.3.132:7003", "192.168.3.132:7004", "192.168.3.132:7005", "192.168.3.132:7006"}, //set redis cluster url
		Addrs: []string{"redis-0.redis.public-services.svc.cluster.local:6379"}, //set redis cluster url
		//Addrs:    []string{"192.168.3.130:6379"}, //set redis cluster url
		//Addrs:    []string{"192.168.3.130:16379", "redis-0.redis.public-services.svc.cluster.local:6379"}, //set redis cluster url
		Password: "", //set password
	})

	return client
}

func main() {
	client := getConnectCluster()
	name, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(name)
	//client.Set(ctx, "name", "LAI", -1)
	//
	//name, err := redis2.String(cluster.Do("get", "name"))
	//fmt.Println(err)
	//fmt.Println(name)
	//client := getConnect()
	////pong, err := client.Ping(ctx).Result()
	//s, err := client.Get(ctx, "name").Result()
	//
	//fmt.Println(s)
	//fmt.Println(err)
	//
	//client := getConnectCluster()
	//s, err := client.Get(ctx, "name").Result()
	//fmt.Println(s, err)
	//
	//getConnect()
	//
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
	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}
