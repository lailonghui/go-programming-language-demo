/*
@Time : 2021/11/22 14:25
@Author : Administrator
@Description :
@File : main
@Software: GoLand
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, f := range files {
			file, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}

	for v, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", v, n)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if strings.EqualFold(input.Text(), "exit") {
			break
		}
		counts[input.Text()]++
	}
}
