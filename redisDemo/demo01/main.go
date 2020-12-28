/*
@Time : 2020/12/25 14:36
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", "192.168.3.130:6379", redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("conn redis failed:", err)
		return
	}

	fmt.Println("redis conn success")

	defer c.Close()

	a, err := redis.String(c.Do("get", "age"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("查询的内容为:%#v", a)

	//reply, err := c.Do("set", "age", 19)
	//if err != nil {
	//	fmt.Println(reply)
	//	return
	//}
	//fmt.Println("插入成功")

}
