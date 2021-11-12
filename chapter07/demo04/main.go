/*
@Time : 2021/8/23 10:41
@Author : Administrator
@Description :
@File : main
@Software: GoLand
*/
package main

import "fmt"

func main() {
	//format := "2006-01-02 15:04:05"
	//sqlUpdatedAt, _ := time.ParseInLocation(format, "2021-08-23 22:00:59", time.Local)
	//
	//currentTime := time.Now()
	//earlyTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 6, 00, 00, 0, currentTime.Location())
	//latterTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 22, 00, 00, 0, currentTime.Location())
	//
	//
	//flag := sqlUpdatedAt.Sub(earlyTime).Seconds() > 0 && sqlUpdatedAt.Sub(latterTime).Seconds() < 0
	//
	//fmt.Println(flag)

	flag1 := true
	flag2 := false
	if flag1 && flag2{
			fmt.Println("flag2")
	} else {
		fmt.Println("no flag1")
	}

}



