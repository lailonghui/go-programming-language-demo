/*
@Time : 2021/3/8 11:06
@Author : lai
@Description :
@File : factory
*/
package main

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	kithttp "github.com/go-kit/kit/transport/http"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func arithmeticFactory(_ context.Context, method, path string) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, e error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}

		tgt, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		tgt.Path = path
		var (
			enc kithttp.EncodeRequestFunc
			dec kithttp.DecodeResponseFunc
		)
		enc, dec = encodeArithmeticRequest, decodeArithmeticResponse
		return kithttp.NewClient(method, tgt, enc, dec).Endpoint(), nil, nil
	}
}

func encodeArithmeticRequest(_ context.Context, req *http.Request, request interface{}) error {
	arithReq := request.(ArithmeticRequest)
	p := "/" + arithReq.RequestType + "/" + strconv.Itoa(arithReq.A) + "/" + strconv.Itoa(arithReq.B)
	req.URL.Path += p
	return nil
}

func decodeArithmeticResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response ArithmeticResponse
	var s map[string]interface{}

	if respCode := resp.StatusCode; respCode >= 400 {
		if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
			return nil, err
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
