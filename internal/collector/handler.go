package collector

import "github.com/labstack/echo/v4"

type Handler interface {
	Collect() echo.HandlerFunc
	RegisterHTTPEndPoints(group *echo.Group)
}
