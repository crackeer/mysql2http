package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Fail(ctx *gin.Context, message string) {
	ctx.PureJSON(http.StatusOK, map[string]interface{}{
		"error": message,
		"data":  nil,
	})
	ctx.Abort()
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.PureJSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  data,
	})
	ctx.Abort()
}
