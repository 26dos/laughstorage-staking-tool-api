package service

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var CustomError = new(customError)

type customError struct{}

func (c *customError) ParameterError(ctx context.Context, msg string) error {
	return gerror.NewCode(gcode.CodeInvalidParameter, msg)
}

func (c *customError) NoData(ctx context.Context, msg string) error {
	return gerror.NewCode(gcode.CodeNotFound, msg)
}

func (c *customError) NoAccess(ctx context.Context, msg string) error {
	return gerror.NewCode(gcode.CodeNotAuthorized, msg)
}

func (c *customError) ServerError(ctx context.Context, msg string) error {
	return gerror.NewCode(gcode.CodeInternalError, msg)
}
