/*
@Time : 2020/11/15 21:13
@Author : lai
@Description :
@File : main
*/
package main

import "fmt"

func main() {
	var f float32 = 1 << 24
	fmt.Println(f == f+1)
}
