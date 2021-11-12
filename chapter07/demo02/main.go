/*
@Time : 2021/7/31 15:06
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"time"
)
func main() {
	format := "2006-01-02 15:04:05"
	sqlUpdatedAt, _ := time.ParseInLocation(format, "2021-07-31 5:01:59", time.Local)
	t := time.Now()
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 2, 0, 0, 0, t.Location())
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 5, 0, 0, 0, t.Location())

	flag := sqlUpdatedAt.Sub(startTime).Milliseconds() > 0 && sqlUpdatedAt.Sub(endTime).Milliseconds() < 0

	fmt.Println(flag)

}
