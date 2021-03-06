/*
@Time : 2021/3/6 11:25
@Author : lai
@Description :
@File : instrument
*/
package instruments

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"golang.org/x/time/rate"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/services"
	"time"
)

var ErrLimitExceed = errors.New("Rate limit exceed! ")

// NewTokenBucketLimiterWithBuildIn 使用x/time/rate 创建限流中间件
func NewTokenBucketLimiterWithBuildIn(bkt *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !bkt.Allow() {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}

// metricMiddleware 定义监控中间件，嵌入Service
// 新增监控指标项：requestCount和requestLatency
type metricMiddleware struct {
	services.Service
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

// Metrics 指标采集方法
func Metrics(requestCount metrics.Counter, requestLatency metrics.Histogram) services.ServiceMiddleware {
	return func(next services.Service) services.Service {
		return metricMiddleware{
			next,
			requestCount,
			requestLatency,
		}
	}
}

func (mw metricMiddleware) Add(a, b int) (ret int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Add"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	ret = mw.Service.Add(a, b)
	return ret
}
