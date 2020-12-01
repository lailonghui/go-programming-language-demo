/*
@Time : 2020/12/1 9:58
@Author : lai
@Description :
@File : main
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	//var x, y IntSet
	//x.Add(1)
	//x.Add(144)
	//x.Add(9)
	//fmt.Println(x.String()) // "{1 9 144}"
	//
	//y.Add(95)
	//y.Add(42)
	//fmt.Println(y.String()) // "{9 42}"
	//
	//x.UnionWith(&y)
	//fmt.Println(x.String())
	//fmt.Println(uint(8 / 64))
	//var a int = 1
	//a |= 2
	//fmt.Println(a)

	var i IntSet
	i.Add(4)
	i.Add(5)
	i.Add(6)
	i.Add(7)
	fmt.Println(i.words)
	//fmt.Println(i.String())
	fmt.Println(i.Has(2))
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	fmt.Println(s.words[word] & (1 << bit))
	fmt.Println(s.words[word] & (1 << bit))
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	fmt.Printf("word=%d,bit=%d\n", word, bit)
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
