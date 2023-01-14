package internal

import (
	"context"
	"github.com/labstack/echo-contrib/prometheus"
)

type Metrics struct{}

func (m *Metrics) ToPrometheusMetrics() []*prometheus.Metric {
	return []*prometheus.Metric{}
}

type metricsContextKey struct{}

func WithMetrics(ctx context.Context, metrics *Metrics) context.Context {
	return context.WithValue(ctx, metricsContextKey{}, metrics)
}

func MustMetrics(ctx context.Context) *Metrics {
	if v := ctx.Value(metricsContextKey{}); v != nil {
		if m, ok := v.(*Metrics); ok {
			return m
		}
	}
	panic("metrics not set on context")
}
