/*
@Time : 2021/2/8 11:17
@Author : lai
@Description :
@File : calculate
*/
package service

import "errors"

// service Define a service interface
type Service interface {
	// Add calculate a+b
	Add(a, b int) int

	// Subtract calculate a-b
	Subtract(a, b int) int

	// Multiply calculate a*b
	Multiply(a, b int) int

	// Divide calculate a/b
	Divide(a, b int) (int, error)
}

// ArithmeticService implement service interface
type ArithmeticService struct {
}

// Add implement Add method
func (as ArithmeticService) Add(a, b int) int {
	return a + b
}

// Subtract implement Subtract method
func (as ArithmeticService) Subtract(a, b int) int {
	return a - b
}

// Multiply implement Multiply method
func (as ArithmeticService) Multiply(a, b int) int {
	return a * b
}

// Divide implement Divide method
func (as ArithmeticService) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("the dividend can not be zero! ")
	}
	return a / b, nil
}
