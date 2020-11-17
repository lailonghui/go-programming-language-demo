/*
@Time : 2020/11/17 14:52
@Author : lai
@Description :
@File : main
*/
package main

import "fmt"

//一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身（该限制同样适用于数组。）
// 但是S类型的结构体可以包含*S指针类型的成员
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
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
	//t1 := &tree{
	//	value: 1,
	//}
	//t2 := add(t1, 2)
	//t2 := appendValues(nil, t1)
	//fmt.Println(t2)
	var a = []int{4, 6, 7, 6, 8, 321, 32, 45, 65, 6, 13, 4565}
	Sort(a)
	fmt.Println(a)

}
