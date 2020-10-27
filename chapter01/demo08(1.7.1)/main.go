/*
@Time : 2020/10/27 10:50
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getURL)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func getURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
