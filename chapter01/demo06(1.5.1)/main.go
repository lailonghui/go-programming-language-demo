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
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		start := time.Now()
		resp, err := http.Get(isContainPrefix(url))
		if err != nil {
			fmt.Printf("get %s failed:%s\n", url, err)
			os.Exit(1)

		}
		nbytes, _ := io.Copy(ioutil.Discard, resp.Body)
		//data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("ioutil.ReadAll() failed:%v ", err)
			os.Exit(1)
		}
		//fmt.Printf("%s",data)
		secs := time.Since(start).Seconds()
		fmt.Printf("%.2fs    %7d    %s\n", secs, nbytes, url)
	}
	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
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
