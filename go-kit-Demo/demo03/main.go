/*
@Time : 2021/3/5 14:31
@Author : lai
@Description :
@File : main
*/
package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/time/rate"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/endpoints"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/instruments"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/services"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/transports"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var svc services.Service
	svc = services.ArithmeticService{}

	// add logging middleware
	svc = services.LoggingMiddleware(logger)(svc)

	//创建指标采集对象：请求次数采集和请求延时采集对象。
	fieldKeys := []string{"method"}
	requestCount := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: "raysonxin",
		Subsystem: "arithmetic_service",
		Name:      "request_count",
		Help:      "Number of requests received",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: "raysonxin",
		Subsystem: "arithmetic_service",
		Name:      "request_latency",
		Help:      "Total duration of requests in microseconds",
	}, fieldKeys)
	//使用Metrics方法对Service对象进行封装：
	svc = instruments.Metrics(requestCount, requestLatency)(svc)

	endpoint := endpoints.MakeArithmeticEndpoint(svc)

	//add ratelimit,refill every second, set capacity 3
	rateBucket := rate.NewLimiter(rate.Every(time.Second), 100)
	endpoint = instruments.NewTokenBucketLimiterWithBuildIn(rateBucket)(endpoint)

	r := transports.MakeHttpHandler(ctx, endpoint, logger)

	go func() {
		fmt.Println("Http Server start at port:9000")
		handler := r
		errChan <- http.ListenAndServe(":9000", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
