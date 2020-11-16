/*
@Time : 2020/11/16 16:12
@Author : lai
@Description :
@File : main
*/
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	data = noneEmpty(data)
	fmt.Println(data)
}

//过滤空字符串
func noneEmpty(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
