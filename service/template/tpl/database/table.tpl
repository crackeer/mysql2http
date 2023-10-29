package {{database}}

import (
    "mysql2http/util"
    "github.com/gin-gonic/gin"
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
    query, setting, err := util.ParseQueryRequest(ctx)
    if err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    sql, values := util.BuildQuery(query)
    db := globalDB.Table("{{table}}").Where(sql, values...)
    if len(setting.Fields) > 0 {
        db = db.Select(setting.Fields)
    }
    if len(setting.OrderBy) > 0 {
        db = db.Order(setting.OrderBy)
    }
    list := []{{table_struct_name}}{}
    err = db.Limit(int(setting.Limit)).Find(&list).Error
    if err != nil {
        util.Fail(ctx, err.Error())
        return
    }

    util.Success(ctx, list)
}

func Create{{table_struct_name}}(ctx *gin.Context)  {
    input := &{{table_struct_name}}{}
    if err := util.ParseCreateRequest(ctx, input); err != nil {
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
    updateData, setting, err := util.ParseModifyRequest(ctx); 
    if err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    if len(setting.Where) < 1 {
        util.Fail(ctx, "no where condition")
        return
    }
    result := globalDB.Model(&{{table_struct_name}}{}).Where(setting.Where).Updates(updateData)
    if result.Error != nil {
        util.Fail(ctx, result.Error.Error())
    } else {
        util.Success(ctx, map[string]interface{}{
            "affected_rows" : result.RowsAffected,
        })
    }    
}


func Delete{{table_struct_name}}(ctx *gin.Context)  {
    query, err := util.ParseDeleteRequest(ctx); 
    if err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    count := globalDB.Where(query).Delete(&{{table_struct_name}}{}).RowsAffected
    util.Success(ctx, map[string]interface{}{
        "affected_rows" : count,
    })
}

func WildQuery{{table_struct_name}}(ctx *gin.Context)  {
     query, setting, err := util.ParseWildQueryRequest(ctx)
    if err != nil {
        util.Fail(ctx, err.Error())
        return
    }
    sql, values := util.BuildQuery(query)
    db := globalDB.Table("{{table}}").Where(sql, values...)
    if len(setting.Fields) > 0 {
        db = db.Select(setting.Fields)
    }
    if len(setting.OrderBy) > 0 {
        db = db.Order(setting.OrderBy)
    }

    if len(setting.GroupBy) > 0 {
        db = db.Group(setting.GroupBy)
    }

    if setting.Offset > 0 {
        db = db.Offset(int(setting.Offset))
    }

     if setting.Limit > 0 {
        db = db.Limit(int(setting.Limit))
    }
    list := []map[string]interface{}{}
    err = db.Find(&list).Error
    if err != nil {
        util.Fail(ctx, err.Error())
        return
    }

    util.Success(ctx, list)
}