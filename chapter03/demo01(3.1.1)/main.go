/*
@Time : 2020/11/14 17:13
@Author : lai
@Description :
@File : main
*/
package main

import "fmt"

func main() {
	//fmt.Println(3 << 5)
	ascci := 'a'
	unicode := '赖'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascci)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d  %[1]q\n", newline)
}
