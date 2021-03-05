/*
@Time : 2021/2/7 15:03
@Author : lai
@Description :
@File : services
*/
package main

import (
	"errors"
	"strings"
)

// StringService provides operations on strings
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
