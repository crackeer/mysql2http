package {{database}}

import (
    "mysql2http/container"
    {% if include_time %}
        "mysql2http/define"
    {% endif %}
)


type {{table_struct_name}} struct {
    {% for item in fields %} 
    {{item.name}}   {{ item.type }} `json:"{{item.field}}" gorm:"{{item.field}}"`
    {% endfor %}
}


func ({{table_struct_name}}) TableName() string {
    return "{{table}}"
}

func init() {
    container.RegisterModelParseFunc("{{database}}", "{{table}}", func() interface{} {
        return &{{table_struct_name}}{}
    })
    container.RegisterModelListParseFunc("{{database}}", "{{table}}", func() interface{} {
        return []{{table_struct_name}}{}
    })
}

