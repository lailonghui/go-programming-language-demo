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
	"os"
)

const URL = "https://docs.hacknode.org/gopl-zh/ch5/ch5-02.html"

func main() {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
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
