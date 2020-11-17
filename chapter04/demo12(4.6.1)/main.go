/*
@Time : 2020/11/18 7:07 AM
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

//var report = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgp": daysAgo}).Parse(templ))

func main() {
	fmt.Println(daysAgo(time.Date(2020, 11, 1, 1, 1, 1, 1, time.Local)))
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
