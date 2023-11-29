package handler

import (
    "github.com/gin-gonic/gin"
    "mysql2http/util"
)

// Query
//
//  @param ctx
func Query(ctx *gin.Context) {
    query, setting, err := util.ParseQueryRequest(ctx)
    if err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    globalDB := getGormDB(ctx)
    sql, values := util.BuildQuery(query)
    db := globalDB.Table(ctx.Param("table")).Where(sql, values...)
    if len(setting.Fields) > 0 {
        db = db.Select(setting.Fields)
    }
    if len(setting.OrderBy) > 0 {
        db = db.Order(setting.OrderBy)
    }
    list := getModelListObject(ctx)
    err = db.Limit(int(setting.Limit)).Find(&list).Error
    if err != nil {
        util.Fail(ctx, err.Error())
        return
    }

    util.Success(ctx, list)
}
