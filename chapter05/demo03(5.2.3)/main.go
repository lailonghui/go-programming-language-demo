/*
@Time : 2020/11/25 15:24
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

func main() {
	resp, _ := http.Get(URL)
	node, _ := html.Parse(resp.Body)
	fmt.Println(len(visit(nil, node)))
	//for _, link := range visit(nil, node) {
	//	fmt.Println(link)
	//}
}

//visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	//fmt.Println(n.Type)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	first := n.FirstChild
	next := n.NextSibling
	if first != nil {
		links = visit(links, first)
	}
	if next != nil {
		links = visit(links, next)
	}

	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	links = visit(links, c)
	//}
	return links
}
