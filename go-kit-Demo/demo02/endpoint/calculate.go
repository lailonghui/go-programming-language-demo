/*
@Time : 2021/2/8 11:31
@Author : lai
@Description :
@File : calculate
*/
package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo02/service"
	"strings"
)

// ArithmeticRequest define request struct
type ArithmeticRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

// ArithmeticResponse define response struct
type ArithmeticResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

//MakeArithmeticEndpoint make endpoint
func MakeArithmeticEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArithmeticRequest)

		var (
			res, a, b int
			calError  error
		)
		a = req.A
		b = req.B
		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		} else if strings.EqualFold(req.RequestType, "Subtract") {
			res = svc.Subtract(a, b)
		} else if strings.EqualFold(req.RequestType, "Multiply") {
			res = svc.Multiply(a, b)
		} else if strings.EqualFold(req.RequestType, "Divide") {
			res, calError = svc.Divide(a, b)
		} else {
			return nil, errors.New("ErrInvalidRequestType")
		}
		return ArithmeticResponse{Result: res, Error: calError}, nil

	}
}
