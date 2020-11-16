/*
@Time : 2020/11/16 9:47
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "hello, world"
	//fmt.Println(len(s))     // "12"
	//fmt.Println(s[11])     // "12"
	//s1 := "赖龙辉a\n\a"
	//fmt.Println(s1)     // "12"
	//fmt.Println(s1[0],s1[1],s1[2],s1[3],s1[4],s1[5],s1[6],s1[7],s1[8],s1[9])     // "12"
	//fmt.Println(utf8.DecodeRuneInString("鲜动"))
	s := "Hello, 世界"
	//for i := 0; i < len(s); {
	//	r, size := utf8.DecodeRuneInString(s[i:])
	//	fmt.Printf("%d\t%c\n", i, r)
	//	i += size
	//}

	//for i,r := range s {
	//	fmt.Printf("%d\t%q\t%d\n",i,r,r)
	//}
	arr := strings.Fields(s)
	fmt.Println(arr)
}
