/*
@Time : 2020/10/23 16:10
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

//Use ctrl+d to stop this program
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		data := input.Text()
		counts[data]++
	}
	for val, num := range counts {
		if num > 1 {
			fmt.Printf("%s:\t%d\n", val, num)
		}
	}
}
