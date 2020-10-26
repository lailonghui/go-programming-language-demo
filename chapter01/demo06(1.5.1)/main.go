/*
@Time : 2020/10/26 10:47
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(isContainPrefix(url))
		if err != nil {
			fmt.Printf("get %s failed:%s\n", url, err)
			os.Exit(1)

		}
		io.Copy(os.Stdout, resp.Body)
		//data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("ioutil.ReadAll() failed:%v ", err)
			os.Exit(1)
		}
		//fmt.Printf("%s",data)

	}
}

func isContainPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") {
		s := strings.Builder{}
		s.WriteString("http://")
		s.WriteString(url)
		return s.String()
	}
	return url
}
