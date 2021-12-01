/*
@Time : 2020/10/20 14:52
@Author : lai
@Description :
@File : main
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

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if strings.EqualFold(input.Text(), "exit") {
			break
		}
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d", line, n)
		}
	}

}
