package server

import (
	"absurdlab.io/tigerd/cmd/server/internal"
	"github.com/labstack/echo/v4"
)

func metricsMiddleware(metrics *internal.Metrics) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := internal.WithMetrics(c.Request().Context(), metrics)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
