package router

import (
    {% for item in databases %}
    "mysql2http/database/{{ item.database }}"
    {% endfor %}
)

func initDatabase() {
    {% for item in databases %}
        if err := {{item.database}}.Init(); err != nil {
            panic(fmt.Errorf("initialize database `%` error: %s", item.database, err))
        }
    {% endfor %}
}


func main() {
    initDatabase()
    router := gin.Default()
    {% for item in databases %}
        {% for table in item.tables %}
            router.POST("/query/{{item.database}}/{{table.table}}", {{item.database}}.Query{{table.name}})
            router.POST("/create/{{item.database}}/{{table.table}}", {{item.database}}.Create{{table.name}})
            router.POST("/modify/{{item.database}}/{{table.table}}", {{item.database}}.Modify{{table.name}})
            router.POST("/delete/{{item.database}}/{{table.table}}", {{item.database}}.Delete{{table.name}})
        {% endfor %}
    {% endfor %}
}