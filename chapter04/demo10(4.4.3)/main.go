/*
@Time : 2020/11/17 15:46
@Author : lai
@Description :
@File : main
*/
package main

import "fmt"

type T struct{ a, b *int }

func main() {

	var a = T{a: new(int)}
	fmt.Println(a)
}
