package handler

import "github.com/labstack/echo/v4"

func (s *statsHandler) RegisterHTTPEndPoints(version *echo.Group) {
	group := version.Group("/stats")
	group.GET("/", s.GetStats())
}
