package handler

import (
	"github.com/gin-gonic/gin"
	"mysql2http/util"
)

func WildQuery(ctx *gin.Context) {
	query, setting, err := util.ParseWildQueryRequest(ctx)
	if err != nil {
		util.Fail(ctx, err.Error())
		return
	}
	sql, values := util.BuildQuery(query)
	globalDB := getGormDB(ctx)
	db := globalDB.Table(ctx.Param("table")).Where(sql, values...)
	if len(setting.Fields) > 0 {
		db = db.Select(setting.Fields)
	}
	if len(setting.OrderBy) > 0 {
		db = db.Order(setting.OrderBy)
	}

	if len(setting.GroupBy) > 0 {
		db = db.Group(setting.GroupBy)
	}

	if setting.Offset > 0 {
		db = db.Offset(int(setting.Offset))
	}

	if setting.Limit > 0 {
		db = db.Limit(int(setting.Limit))
	}
	list := []map[string]interface{}{}
	err = db.Find(&list).Error
	if err != nil {
		util.Fail(ctx, err.Error())
		return
	}

	util.Success(ctx, list)
}
