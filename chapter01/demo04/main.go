/*
@Time : 2020/10/26 9:04
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, v := range strings.Split(string(data), "\n") {
			counts[v]++
		}
	}
	for key, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, key)
		}
	}
}
