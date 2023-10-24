package {{database}}

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)
{% autoescape off %}
var dsn string = "{{dsn}}"
var globalDB *gorm.DB
{% endautoescape %}
func init() {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
    globalDB = db
}