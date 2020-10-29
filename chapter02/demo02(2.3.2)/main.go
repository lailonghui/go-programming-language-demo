/*
@Time : 2020/10/29 14:49 
@Author : lai
@Description :
@File : main
*/
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n1",false,"omit trailing newline")
var sep = flag.String("s1"," ","separator")
func main()  {
	flag.Parse()
	fmt.Println(*n)
	fmt.Println(*sep)
	fmt.Println(flag.Args())
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n{
		fmt.Println()
	}
}
