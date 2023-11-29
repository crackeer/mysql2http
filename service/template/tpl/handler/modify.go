package handler

import (
	"github.com/gin-gonic/gin"
	"mysql2http/util"
)

func Modify(ctx *gin.Context) {
    updateData, setting, err := util.ParseModifyRequest(ctx); 
    if err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    if len(setting.Where) < 1 {
        util.Fail(ctx, "no where condition")
        return
    }
    globalDB := getGormDB(ctx)
    result := globalDB.Table(ctx.Param("table")).Where(setting.Where).Updates(updateData)
    if result.Error != nil {
        util.Fail(ctx, result.Error.Error())
    } else {
        util.Success(ctx, map[string]interface{}{
            "affected_rows" : result.RowsAffected,
        })
    }    
}