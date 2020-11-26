/*
@Time : 2020/11/26 10:21
@Author : lai
@Description :
@File : main
*/
package main

import "log"

func main() {
	log.SetPrefix("wait:")
	log.SetFlags(0)
	log.Fatalf("Site is down: %v\n", "aa")
}
