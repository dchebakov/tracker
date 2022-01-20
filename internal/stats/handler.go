package stats

import "github.com/labstack/echo/v4"

type Handler interface {
	GetStats() echo.HandlerFunc
	RegisterHTTPEndPoints(version *echo.Group)
}
