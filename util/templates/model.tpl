package {{database}}

import (
    "gorm.io/gorm"
)

type {{table_struct_name}} struct {
    {% for item in fields %} 
    {{item.name}}   {{ item.type }} `json:"{{item.field}}" gorm:"{{item.field}}"`
    {% endfor %}
}

func ({{table_struct_name}}) TableName() string {
    return "{{table}}"
}

func Query(db *gorm.DB, query map[string]interface{}, limit int64) ([]{{table_struct_name}}, error) {
    list := []{{table_struct_name}}{}
    err := db.Model(&{{table_struct_name}}{}).Where(query).Limit(int(limit)).Find(&list).Error
    return list, err
}

func Query(db *gorm.DB, query map[string]interface{}, limit int64) ([]{{table_struct_name}}, error) {
    list := []{{table_struct_name}}{}
    err := db.Model(&{{table_struct_name}}{}).Where(query).Limit(int(limit)).Find(&list).Error
    return list, err
}