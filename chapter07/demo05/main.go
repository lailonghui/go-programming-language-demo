/*
@Time : 2021/11/17 20:40
@Author : Administrator
@Description :
@File : main
@Software: GoLand
*/
package main

import "fmt"

func main() {
	//i := "24"
	var i = 1
	//转4位数字符串，不足前补0
	newStr := fmt.Sprintf("%03d", i)
	fmt.Println(newStr) // 0001

}
