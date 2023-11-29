package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mysql2http/container"
)

func getGormDB(ctx *gin.Context) *gorm.DB {
	database := ctx.Param("database")
	return container.GetGormDB(database)
}

func getModelObject(ctx *gin.Context) interface{} {
	database := ctx.Param("database")
	table := ctx.Param("table")
	return container.GetModelObject(database, table)
}

func getModelListObject(ctx *gin.Context) interface{} {
	database := ctx.Param("database")
	table := ctx.Param("table")
	return container.GetModelListObject(database, table)
}
