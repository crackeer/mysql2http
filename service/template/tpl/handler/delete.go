package handler

import (
    "github.com/gin-gonic/gin"
    "mysql2http/util"
)

func Delete(ctx *gin.Context) {
    input := getModelObject(ctx)
    globalDB := getGormDB(ctx)

    query, err := util.ParseDeleteRequest(ctx)
    if err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    count := globalDB.Where(query).Delete(input).RowsAffected
    util.Success(ctx, map[string]interface{}{
        "affected_rows": count,
    })
}
