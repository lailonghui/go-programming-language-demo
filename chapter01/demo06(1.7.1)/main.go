/*
@Time : 2021/11/22 15:49
@Author : Administrator
@Description :
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":10086", nil)
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
