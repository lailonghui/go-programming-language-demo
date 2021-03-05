/*
@Time : 2021/2/9 16:14
@Author : lai
@Description :
@File : loggings
*/
package service

import (
	"github.com/go-kit/kit/log"
	"time"
)

// MiddlewareService define services middleware
type MiddlewareService func(Service) Service

// loggingMiddleware Make a new type
// that contains Service interface and logger instance
type loggingMiddleware struct {
	Service
	logger log.Logger
}

// LoggingMiddleware make logging middleware
func LoggingMiddleware(logger log.Logger) MiddlewareService {
	return func(next Service) Service {
		return loggingMiddleware{next, logger}
	}
}

func (mw loggingMiddleware) Add(a, b int) (ret int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Add",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(begin),
		)
	}(time.Now())

	ret = mw.Service.Add(a, b)
	return ret
}
