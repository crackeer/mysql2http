package main

import (
    _ "mysql2http/container"
    "mysql2http/handler"
    "github.com/gin-gonic/gin"
    "os"
    "strconv"
    {% for item in databases %}
    _ "mysql2http/database/{{ item.database }}"
    {% endfor %}
)

var (
    port string = "{{port}}"
)


func main() {
    if value, exist := os.LookupEnv("PORT"); exist && len(value) > 0 {
        if _, err := strconv.Atoi(value); err != nil {
            port = value
        }
    }
    router := gin.New()
	gin.SetMode(gin.DebugMode)
    router.POST("/:database/:table/query", handler.Query)
    router.POST("/:database/:table/wild_query", handler.WildQuery)
    router.POST("/:database/:table/create", handler.Create)
    router.POST("/:database/:table/modify", handler.Modify)
    router.POST("/:database/:table/delete", handler.Delete)
    router.GET("/:database/:table/fields", handler.Fields)
    router.NoRoute(handler.Default)
    router.Run(":" + port)
}