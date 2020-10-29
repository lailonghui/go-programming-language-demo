/*
@Time : 2020/10/29 14:22 
@Author : lai
@Description :
@File : main
*/
package main

import "fmt"

func main()  {
	x := 1
	p := &x
	fmt.Println(p)
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
}
