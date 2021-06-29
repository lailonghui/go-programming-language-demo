/*
@Time : 2021/6/22 9:54
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"github.com/xingliuhua/leaf"
	"time"
)

func main() {
	err, node := leaf.NewNode(0)
	if err != nil {
		return
	}
	// 每毫秒200个
	err = node.SetGenerateIDRate(200)
	if err != nil {
		return
	}
	startTime := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000
	// 从2020年开始
	err = node.SetSince(startTime)
	if err != nil {
		return
	}
	err, id := node.NextId()

	fmt.Println(id)
}
