/*
@Time : 2020/11/30 14:49
@Author : lai
@Description :
@File : main
*/
package main

import "os"

func main() {
	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d               // NOTE: necessary!
		os.MkdirAll(dir, 0755) // creates parent directories too
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
	// ...do some work…
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}

func tempDirs() []string {
	return []string{"./a", "./b"}
}
