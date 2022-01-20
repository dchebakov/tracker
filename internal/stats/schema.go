package stats

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Filter struct {
	CustomerID *int64
	Day        *time.Time
}

func (f *Filter) Bind(c echo.Context) error {
	params := c.QueryParams()
	if params.Has("customerID") {
		f.CustomerID = new(int64)
	}
	if params.Has("day") {
		f.Day = new(time.Time)
	}

	return echo.QueryParamsBinder(c).
		Int64("customerID", f.CustomerID).
		Time("day", f.Day, "2006-01-02").BindError()
}
