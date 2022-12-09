package middleware

import (
	"net/http"

	"github.com/frtatmaca/case/errors/httperror"
	"github.com/frtatmaca/case/reqctx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandler(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			logger.Errorw(c.Request.URL.Path,
				zap.Any("error", err),
				zap.String(reqctx.CorrelationIdHeader, reqctx.GetCorrelationId(c)))
			switch e := err.Err.(type) {
			case httperror.HttpError:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Service Unavailable"})
			}
		}
	}
}
