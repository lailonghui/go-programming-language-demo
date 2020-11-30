/*
@Time : 2020/11/30 14:57
@Author : lai
@Description :
@File : main
*/
package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...))
}

//sum return the sum of any number of int
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
