/*
@Time : 2020/11/16 17:05
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	for name := range ages {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
