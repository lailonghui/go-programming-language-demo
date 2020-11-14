/*
@Time : 2020/11/13 16:25
@Author : lai
@Description :
@File : main
*/
package main

import "fmt"

func main() {
	var a interface{}
	a = "3424"
	t, ok := a.(int)

	fmt.Println(t, ok)
}

//斐波纳契数列
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

//最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
