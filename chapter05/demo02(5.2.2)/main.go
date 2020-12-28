/*
@Time : 2020/11/25 11:15
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
	//for _, link := range visit(nil, node) {
	//	fmt.Println(link)
	//}
	outline(nil, node)
}

//outline output of the tree structure
func outline(stack []string, n *html.Node) {
	fmt.Println(n.Type)
	//if n.Type == html.ElementNode {
	//	stack = append(stack, n.Data)
	//	fmt.Println(stack)
	//}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}

}

//visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
