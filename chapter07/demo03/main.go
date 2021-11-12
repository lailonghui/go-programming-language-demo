package main

import "fmt"

func main()  {
	score:= 1.222213
	a := fmt.Sprintf("%.6f:%.6f", score,score)
	fmt.Println(a)

}