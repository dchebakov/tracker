package health

import "github.com/labstack/echo/v4"

type Handler interface {
	Health() echo.HandlerFunc
	Readiness() echo.HandlerFunc
	RegisterHTTPEndPoints(group *echo.Group)
}
