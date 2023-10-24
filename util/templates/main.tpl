package main

import (
    {% for item in databases %}
    "mysql2http/database/{{ item.database }}"
    {% endfor %}
    "github.com/gin-gonic/gin"
)


func main() {
    router := gin.New()
	gin.SetMode(gin.DebugMode)
    {% for item in databases %}
        {% for table in item.tables %}
            router.POST("/query/{{item.database}}/{{table.table}}", {{item.database}}.Query{{table.table_struct_name}})
            router.POST("/create/{{item.database}}/{{table.table}}", {{item.database}}.Create{{table.table_struct_name}})
            //router.POST("/modify/{{item.database}}/{{table.table}}", {{item.database}}.Modify{{table.table_struct_name}})
            //router.POST("/delete/{{item.database}}/{{table.table}}", {{item.database}}.Delete{{table.table_struct_name}})
        {% endfor %}
    {% endfor %}
    router.Run(":8080")
}