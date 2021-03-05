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
	"lai.com/go_programming_language_demo/go-kit-Demo/demo02/service"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/endpoints"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/services"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/transports"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	errChan := make(chan error)

	var svc service.Service
	svc = services.ArithmeticService{}
	endpoint := endpoints.MakeArithmeticEndpoint(svc)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

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
