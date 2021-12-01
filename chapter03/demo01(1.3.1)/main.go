/*
@Time : 2021/11/25 15:25
@Author : Administrator
@Description :
@File : main
@Software: GoLand
*/
package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x) // 0010 0010
	fmt.Printf("%08b\n", y) // 0000 0110
	fmt.Println("---------")
	fmt.Println("---------")
	fmt.Printf("%08b\n", x&y)  // 0000 0010
	fmt.Printf("%08b\n", x|y)  // 0010 0110
	fmt.Printf("%08b\n", x^y)  // 0010 0100
	fmt.Printf("%08b\n", x&^y) // 0010 0000
	fmt.Printf("%08b\n", y&^x) // 0000 0100
	fmt.Println("---------")
	fmt.Println("---------")
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Println("---------")
	fmt.Println("---------")
	fmt.Printf("%08b\n", x<<1) // 0100 0100
	fmt.Printf("%08b\n", x>>1) // 0001 0001
	fmt.Println("---------")
	fmt.Println("---------")
	f := 1e100  // a float64
	i := int(f) // 结果依赖于具体实现
	fmt.Println(i)
}
