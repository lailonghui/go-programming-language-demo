/*
@Time : 2020/11/16 17:23
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
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if !seen[input.Text()] {
			seen[input.Text()] = true
			fmt.Println(input.Text())
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
