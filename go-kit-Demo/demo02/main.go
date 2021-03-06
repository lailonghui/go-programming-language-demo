/*
@Time : 2021/2/9 14:57
@Author : lai
@Description :
@File : main
*/
package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo02/endpoint"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo02/service"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo02/transport"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	var svc service.Service
	svc = service.ArithmeticService{}

	// add logging middleware
	svc = service.LoggingMiddleware(logger)(svc)

	endpoint := endpoint.MakeArithmeticEndpoint(svc)

	r := transport.MakeHttpHandler(ctx, endpoint, logger)

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
