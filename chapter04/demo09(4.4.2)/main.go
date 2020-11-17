/*
@Time : 2020/11/17 15:20
@Author : lai
@Description :
@File : main
*/
package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

//Sort sorts values from small to large
func Sort(values []int) {
	var root *tree
	for _, value := range values {
		root = add(root, value)
	}
	appendValues(values[:0], root)
}

//appendValues appends the elements of t to values in order
//from left to value to right
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

//Add value to target tree. if value greater than t.value,create a new tree to t.left, else create a new tree to t.right
func add(t *tree, value int) *tree {
	if t == nil {
		t = &tree{
			value: value,
		}
		return t
	}
	if value > t.value {
		t.right = add(t.right, value)
	} else {
		t.left = add(t.left, value)
	}
	return t
}

func main() {
	//t1 := &tree{
	//	value: 111,
	//}
	//var t1 *tree
	//t2 := add(t1, 24)
	//fmt.Printf("%#v\n", t2)
	//t1 := &tree{
	//	value: 1,
	//	left: &tree{
	//		value: 2,
	//		left: &tree{
	//			value: 4,
	//		},
	//		right: &tree{
	//			value: 5,
	//		},
	//	},
	//	right: &tree{
	//		value: 3,
	//		left: &tree{
	//			value: 6,
	//		},
	//		right: &tree{
	//			value: 7,
	//		},
	//	},
	//}
	//t2 := appendValues([]int{}, t1)
	//fmt.Println(t2)
	var a = []int{4, 6, 7, 6, 8, 321, 32, 45, 65, 6, 13, 4565}
	Sort(a)
	fmt.Println(a)
}
