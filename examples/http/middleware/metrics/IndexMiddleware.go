package metrics

import (
	"context"
	"github.com/go-kit/kit/metrics"
	"github.com/xinlianit/kit-scaffold/examples/http/endpoint"
	"time"
)

// IndexMiddleware Index中间件类型
type IndexMiddleware func(indexEndpoint endpoint.IndexEndpoint) IndexMetrics

// IndexMetrics 监控指标
type IndexMetrics struct {
	next endpoint.IndexEndpoint
	counter   metrics.Counter
	histogram metrics.Histogram
}

// IndexMetricsMiddleware Index 监控中间件
func IndexMetricsMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) IndexMiddleware {
	return func(next endpoint.IndexEndpoint) IndexMetrics {
		return IndexMetrics{next: next, counter: requestCount, histogram: requestLatency}
	}
}

// Hello 监控指标
func (m IndexMetrics) Hello(ctx context.Context, request interface{}) (response interface{}, err error)  {
	// 监控埋点
	defer func(begin time.Time) {
		labelValues := []string{"method", "Hello"}
		m.counter.With(labelValues...).Add(1)
		m.histogram.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return m.next.Hello(ctx, request)
}

// Test 监控指标
func (m IndexMetrics) Test(ctx context.Context, request interface{}) (response interface{}, err error)  {
	// 监控埋点
	defer func(begin time.Time) {
		labelValues := []string{"method", "Test"}
		m.counter.With(labelValues...).Add(1)
		m.histogram.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return m.next.Test(ctx, request)
}