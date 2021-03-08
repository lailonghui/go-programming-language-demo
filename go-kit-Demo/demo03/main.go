/*
@Time : 2021/3/5 14:31
@Author : lai
@Description :
@File : main
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/time/rate"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/endpoints"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/instruments"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/register"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/services"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/transports"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//定义环境变量
	var (
		consulHost  = flag.String("consul.host", "", "consul ip address")
		consulPort  = flag.String("consul.port", "", "consul port")
		serviceHost = flag.String("service.host", "", "service ip address")
		servicePort = flag.String("service.port", "", "service port")
	)
	//parse
	flag.Parse()

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

	// 创建健康检查的Endpoint,未增加限流
	healthEndpoint := endpoints.MakeHealthCheckEndpoint(svc)

	//把算术运算Endpoint和健康检查Endpoint封装至ArithmeticEndpoints
	endpts := endpoints.ArithmeticEndpoints{
		ArithmeticEndpoint:  endpoint,
		HealthCheckEndpoint: healthEndpoint,
	}

	r := transports.MakeHttpHandler(ctx, endpts, logger)

	//创建注册对象
	registrar := register.Register(*consulHost, *consulPort, *serviceHost, *servicePort, logger)

	go func() {
		fmt.Println("Http Server start at port:" + *servicePort)
		//启动前执行注册
		registrar.Register()
		handler := r
		errChan <- http.ListenAndServe(":"+*servicePort, handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	//服务退出取消注册
	registrar.Deregister()
	fmt.Println(error)
}
