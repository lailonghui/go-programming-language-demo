/*
@Time : 2021/3/5 15:16
@Author : lai
@Description :
@File : services
*/
package services

import "errors"

// Service Define a services interface
type Service interface {

	//Add calculate a+b
	Add(a, b int) int

	// Subtract calculate a-b
	Subtract(a, b int) int

	// Multiply calculate a*b
	Multiply(a, b int) int

	// Divide calculate a/b
	Divide(a, b int) (int, error)

	// HealthCheck check service health status
	HealthCheck() bool
}

// ArithmeticService implement Service interface
type ArithmeticService struct {
}

// ArithmeticService实现HealthCheck
// HealthCheck implement Service method
// 用于检查服务的健康状态，这里仅仅返回true
func (s ArithmeticService) HealthCheck() bool {
	return true
}

// Add implement Add method
func (s ArithmeticService) Add(a, b int) int {
	return a + b
}

// Subtract implement Subtract method
func (s ArithmeticService) Subtract(a, b int) int {
	return a - b
}

// Multiply implement Multiply method
func (s ArithmeticService) Multiply(a, b int) int {
	return a * b
}

func (s ArithmeticService) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("the divided can not be zero! ")
	}
	return a / b, nil
}
