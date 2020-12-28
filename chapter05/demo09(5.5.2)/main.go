/*
@Time : 2020/11/26 14:50
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

const URL = "https://docs.hacknode.org/gopl-zh/ch5/ch5-02.html"

func main() {
	resp, _ := http.Get(URL)
	node, _ := html.Parse(resp.Body)
	forEachNode(node, startElement, endElement)
}

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

//startElement is called before the node
func startElement(n *html.Node) {
	var attrString strings.Builder
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			attrString.WriteString(fmt.Sprintf("\t%s=%q", attr.Key, attr.Val))
		}
		if n.Data == "img" {
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attrString.String())
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attrString.String())
		}
		depth++
	}
}

//endElement is called after the node
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.Data == "img" || n.Data == "meta" || n.Data == "link" {
			return
		}

		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
