/*
@Time : 2020/11/26 15:32
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
	if n := ElementByClassName(node, "chapter ", look); n != nil {
		fmt.Printf("<%s></%s>", n.Data, n.Data)
	}

}

var node *html.Node
var count int

func ElementByClassName(n *html.Node, className string, pre func(n *html.Node, className string) bool) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if ok := look(n, className); !ok {
			node = n
			break
		}
		ElementByClassName(c, className, pre)

	}
	return node
}

func look(n *html.Node, className string) bool {
	for _, a := range n.Attr {
		if a.Key == "class" && a.Val == className {
			return false
		}
	}
	return true
}
