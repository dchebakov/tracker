package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ReadRequest(ctx echo.Context, request interface{}, validate *validator.Validate) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}

	return validate.StructCtx(ctx.Request().Context(), request)
}
