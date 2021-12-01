/*
@Time : 2021/11/22 14:25
@Author : Administrator
@Description :
@File : main
@Software: GoLand
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
	files := os.Args[1:]
	for _, fileName := range files {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err=%v\n", err)
			continue
		}
		for _, s := range strings.Split(string(data), "\n") {
			counts[s]++
		}
	}
	for v, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, v)
		}
	}

}
