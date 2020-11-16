/*
@Time : 2020/11/16 16:40
@Author : lai
@Description :
@File : main
*/
package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
	fmt.Println(s)
}

func remove(dst []int, idx int) []int {
	copy(dst[idx:], dst[idx+1:])
	return dst[:len(dst)-1]
}
