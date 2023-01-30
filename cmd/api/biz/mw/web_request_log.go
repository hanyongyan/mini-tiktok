package mw

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func RequestLogger() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		uri := c.Request.RequestURI()
		method := c.Request.Method()
		hlog.Infof("method: %s\turi: %s", string(method), string(uri))
		c.Next(ctx)
	}
}
