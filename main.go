package main

import (
	"fmt"

	"github.com/crackeer/mysql2http/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:aaa@tcp(calcnode.site:3306)/vrapi?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	tables := util.Tables(db)

	for _, table := range tables {
		fields, _ := util.Desc(db, table)
		bytes, err := util.GenerateModel("vrapi", table, fields)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
		break
	}
}
