package {{database}}

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)
var dsn string = {{dsn}}
var globalDB *gorm.DB

func Init() error {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
    globalDB = db
}