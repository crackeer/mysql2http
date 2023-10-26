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
            router.POST("/{{item.database}}/{{table.table}}/query", {{item.database}}.Query{{table.table_struct_name}})
            router.POST("/{{item.database}}/{{table.table}}/wild_query", {{item.database}}.WildQuery{{table.table_struct_name}})
            router.POST("/{{item.database}}/{{table.table}}/create", {{item.database}}.Create{{table.table_struct_name}})
            router.POST("/{{item.database}}/{{table.table}}/modify", {{item.database}}.Modify{{table.table_struct_name}})
            router.POST("/{{item.database}}/{{table.table}}/delete", {{item.database}}.Delete{{table.table_struct_name}})
        {% endfor %}
    {% endfor %}
    router.Run(":8080")
}