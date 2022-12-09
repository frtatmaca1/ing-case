package reqctx

import (
	"context"
	"fmt"
)

const (
	CorrelationIdHeader = "x-correlationid"
)

func getStringHeader(ctx context.Context, name string) string {
	value := ctx.Value(name)
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}

func GetCorrelationId(ctx context.Context) string {
	return getStringHeader(ctx, CorrelationIdHeader)
}
