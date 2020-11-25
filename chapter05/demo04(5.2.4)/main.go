/*
@Time : 2020/11/25 16:43
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

const URL = "https://docs.hacknode.org/gopl-zh/ch5/ch5-02.html"

var countMap = make(map[string]int)

func main() {
	resp, _ := http.Get(URL)
	node, _ := html.Parse(resp.Body)
	count(node)
	num := 0
	for key, val := range countMap {
		num += val
		fmt.Printf("%s:\t%d\n", key, val)
	}
	fmt.Printf("total number of elements:\t%d", num)
}

//count record the number of occurrences of the element with the same name in the HTML thee
func count(n *html.Node) {
	if n.Type == html.ElementNode {
		if _, ok := countMap[n.Data]; ok {
			countMap[n.Data]++
		} else {
			countMap[n.Data] = 1
		}
	}
	first := n.FirstChild
	next := n.NextSibling
	if first != nil {
		count(first)
	}
	if next != nil {
		count(next)
	}
}
