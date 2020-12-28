/*
@Time : 2020/11/26 9:19
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
	outputContent(node)
}

//outputContent output the content of all text nodes
func outputContent(n *html.Node) {
	//fmt.Println(n.DataAtom)
	if n.Type == html.TextNode && n.Data != "" {
		fmt.Println(n)
	}
	first := n.FirstChild
	next := n.NextSibling
	if first != nil {
		outputContent(first)
	}
	if next != nil {
		outputContent(next)
	}
}
