package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type middlewareManager struct {
	logger *zap.SugaredLogger
}

func NewMiddlewareManager(logger *zap.SugaredLogger) *middlewareManager {
	return &middlewareManager{logger: logger}
}

func (mw *middlewareManager) RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		start := time.Now()

		err := next(ctx)

		req := ctx.Request()
		s := time.Since(start).String()

		msg := "Calling API: " + req.URL.String()
		mw.logger.Infow(msg, "Method", req.Method, "Time", s)
		return err
	}
}
