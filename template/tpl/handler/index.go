package handler

import (
	"embed"
	"io/fs"
	"mysql2http/container"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

var (
	//go:embed static/*
	staticFs embed.FS
)

func Default(ctx *gin.Context) {
	subFS, _ := fs.Sub(staticFs, "static")
	fileServer := http.FileServer(http.FS(subFS))
	fileServer.ServeHTTP(ctx.Writer, ctx.Request)
}
