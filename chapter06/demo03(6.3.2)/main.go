/*
@Time : 2020/12/1 9:03
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func lookUp(key string) string {
	cache.Lock()
	result := cache.mapping[key]
	cache.Unlock()
	return result
}
func main() {
	cache.mapping["name"] = "赖龙辉"
	fmt.Println(lookUp("name"))
}
