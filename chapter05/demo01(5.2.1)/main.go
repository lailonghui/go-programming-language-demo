/*
@Time : 2020/11/24 14:33
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

//const URL = "http://mall.ushirt.info/profit#/partnerCenter"

func main() {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	doc, err := html.Parse(resp.Body)
	//fmt.Println(doc.FirstChild)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	//	os.Exit(1)
	//}
	//for _, link := range visit(nil, doc) {
	//	a = link
	//}
	//fmt.Println(visit(nil, doc))
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	fmt.Println(n.FirstChild)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
