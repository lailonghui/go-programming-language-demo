/*
@Time : 2020/10/29 15:13
@Author : lai
@Description :
@File : main
*/
package main

import "fmt"

func main() {
	p := new(int)
	*p = 12
	fmt.Println(p)
	fmt.Println(*p)
}
