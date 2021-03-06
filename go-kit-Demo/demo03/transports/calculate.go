/*
@Time : 2021/3/5 15:42
@Author : lai
@Description :
@File : calculate
*/
package transports

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/endpoints"
	"net/http"
	"strconv"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

//解码器：把用户的请求内容转换为请求对象（ArithmeticRequest）
// decodeArithmeticRequest decode request params to struct
func decodeArithmeticRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	requestType, ok := vars["type"]

	if !ok {
		return nil, ErrorBadRequest
	}

	pa, ok := vars["a"]

	if !ok {
		return nil, ErrorBadRequest
	}

	pb, ok := vars["b"]
	if !ok {
		return nil, ErrorBadRequest
	}

	a, _ := strconv.Atoi(pa)
	b, _ := strconv.Atoi(pb)

	return endpoints.ArithmeticRequest{
		RequestType: requestType,
		A:           a,
		B:           b,
	}, nil

}

//编码器：把处理结果转换为响应对象（ArithmeticResponse）
// encodeArithmeticResponse encode response to return
func encodeArithmeticResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// MakeHttpHandler make http handler use mux
func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
	}

	//新增用于Prometheus轮询拉取监控指标的代码，开发API接口/metrics
	r.Path("/metrics").Handler(promhttp.Handler())

	r.Methods("POST").Path("/calculate/{type}/{a}/{b}").Handler(kithttp.NewServer(
		endpoint,
		decodeArithmeticRequest,
		encodeArithmeticResponse,
		options...,
	))

	return r
}
