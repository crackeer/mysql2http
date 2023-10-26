package {{database}}

import (
    "mysql2http/util"
    "github.com/gin-gonic/gin"
    "mysql2http/define"
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

func Query{{table_struct_name}}(ctx *gin.Context) {
    input := &define.QueryRequest{}
    if err := ctx.BindJSON(input); err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    if input.PageSize < 1 {
        input.PageSize = 10
    }
    sql, values := util.BuildQuery(input.Query)
    var (
        list []{{table_struct_name}}
        total int64
    )
    err := globalDB.Model(&{{table_struct_name}}{}).Where(sql, values...).Offset(int(util.Offset(input.Page, input.PageSize))).Limit(int(input.PageSize)).Find(&list).Error
    if err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    globalDB.Model(&{{table_struct_name}}{}).Where(sql, values...).Count(&total)

    util.Success(ctx, map[string]interface{}{
        "list" : list,
        "page" : input.Page,
        "page_size" : input.PageSize,
        "total" : total,
        "total_page" : util.TotalPage(total, input.PageSize),
    })
}

func Create{{table_struct_name}}(ctx *gin.Context)  {
    input := &{{table_struct_name}}{}
    if err := ctx.BindJSON(input); err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    if err := globalDB.Create(input).Error; err != nil {
        util.Fail(ctx,  err.Error())
        return
    }
    util.Success(ctx, input)
}

func Modify{{table_struct_name}}(ctx *gin.Context)  {
    input := &define.QueryRequest{}
    if err := ctx.BindJSON(input); err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    sql, values := util.BuildQuery(input.Query)
    count := globalDB.Model(&{{table_struct_name}}{}).Where(sql, values...).Updates(input.Modify).RowsAffected
    util.Success(ctx, map[string]interface{}{
        "affected_rows" : count,
    })
}

func Delete{{table_struct_name}}(ctx *gin.Context)  {
    input := &define.QueryRequest{}
    if err := ctx.BindJSON(input); err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    sql, values := util.BuildQuery(input.Query)
    count := globalDB.Where(sql, values...).Delete(&{{table_struct_name}}{}).RowsAffected
    util.Success(ctx, map[string]interface{}{
        "affected_rows" : count,
    })
}

func Delete{{table_struct_name}}(ctx *gin.Context)  {
    input := &define.QueryRequest{}
    if err := ctx.BindJSON(input); err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    sql, values := util.BuildQuery(input.Query)
    count := globalDB.Where(sql, values...).Delete(&{{table_struct_name}}{}).RowsAffected
    util.Success(ctx, map[string]interface{}{
        "affected_rows" : count,
    })
}
