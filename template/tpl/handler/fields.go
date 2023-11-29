package handler

import (
    "github.com/gin-gonic/gin"
    "mysql2http/util"
)

// Query
//
//  @param ctx
func Fields(ctx *gin.Context) {
    util.Success(ctx, []map[string]interface{}{})
}
