/*
@Time : 2020/10/23 16:36
@Author : lai
@Description :
@File : main
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, name := range files {
			file, err := os.Open(name)
			if err != nil {
				fmt.Printf("os.Open() err:%v", err)
				return
			}
			countLines(file, counts)
		}
	}
	for val, line := range counts {
		fmt.Printf("%s:\t%d\n", val, line)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		data := input.Text()
		if data == "exit" {
			return
		}
		counts[data]++
	}
}
