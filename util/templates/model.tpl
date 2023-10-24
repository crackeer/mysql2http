package {{database}}

import (
    "gorm.io/gorm"
    "mysql2http/util"
)


type {{table_struct_name}} struct {
    {% for item in fields %} 
    {{item.name}}   {{ item.type }} `json:"{{item.field}}" gorm:"{{item.field}}"`
    {% endfor %}
}

func ({{table_struct_name}}) TableName() string {
    return "{{table}}"
}

func Query{{table_struct_name}}(ctx *gin.Context) {
    input := &define.QueryRequest{}
    if err := ctx.BindJSON(input); err != nil {
        util.Fail(ctx, err)
        return
    }
    sql, values := util.BuildQuery(input.Query)
    var (
        list []{{table_struct_name}}
        total int64
    )
    err := globalDB.Model(&{{table_struct_name}}{}).Where(sql, values...).Offset(int(util.Offset(input.Page, input.PageSize))).Limit(int(input.PageSize)).Find(&list).Error
    if err != nil {
        util.Fail(ctx, err)
        return
    }
    globalDB.Model(&{{table_struct_name}}{}).Where(sql, values...).Count(&total)

    util.Success(ctx, map[string]interface{}{
        "list" : list,
        "page" : input.Page,
        "page_size" : input.PageSize,
        "total" : util.TotalPage(total, input.PageSize),
        "total_page" : input.TotalSize,
    })
}

func Create{{table_struct_name}}(ctx *gin.Context, query map[string]interface{}, limit int64) ([]{{table_struct_name}}, error) {
    input := &{{table_struct_name}}{}
    if err := ctx.BindJSON(input); err != nil {
        util.Fail(ctx, err)
        return
    }
    if err := db.Create(input).Error; err != nil {
        util.Fail(ctx, err)
        return
    }
    util.Success(ctx, input)
}
