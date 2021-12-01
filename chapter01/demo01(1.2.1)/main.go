/*
@Time : 2020/10/20 14:52
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	var s, sep string
	sep = " "

	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + sep
	}
	fmt.Println(s)
}
