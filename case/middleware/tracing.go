package middleware

import (
	"github.com/frtatmaca/case/reqctx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TracingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationID := c.Request.Header.Get(reqctx.CorrelationIdHeader)
		if correlationID == "" {
			correlationID = uuid.New().String()
		}
		c.Set(reqctx.CorrelationIdHeader, correlationID)
		c.Next()
	}
}
