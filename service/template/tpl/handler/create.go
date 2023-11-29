package handler

import (
	"github.com/gin-gonic/gin"
	"mysql2http/util"
)

func Create(ctx *gin.Context) {
	input := getModelObject(ctx)
	globalDB := getGormDB(ctx)
	if err := util.ParseCreateRequest(ctx, input); err != nil {
		util.Fail(ctx, err.Error())
		return
	}
	if err := globalDB.Create(input).Error; err != nil {
		util.Fail(ctx, err.Error())
		return
	}
	util.Success(ctx, input)
}
