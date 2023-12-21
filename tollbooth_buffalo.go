package tollbooth_buffalo

import (
	"fmt"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/gobuffalo/buffalo"
)

func LimitHandler(lmt *limiter.Limiter) buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			httpError := tollbooth.LimitByRequest(lmt, c.Response(), c.Request())
			if httpError != nil {
				return c.Error(httpError.StatusCode, fmt.Errorf("%s", httpError.Message))
			}

			return next(c)
		}
	}
}
