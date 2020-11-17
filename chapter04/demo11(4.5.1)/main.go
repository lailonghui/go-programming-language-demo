/*
@Time : 2020/11/17 16:08
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"net/http"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	//data, _ := json.MarshalIndent(movies, "", "  ")
	//fmt.Println(string(data))
	http.HandleFunc("/", query)
	http.ListenAndServe(":8080", nil)
}

func query(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
	fmt.Println(r.Host)
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.URL)
}
