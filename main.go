package main

import (
	"fmt"

	"github.com/crackeer/mysql2http/util"
)

func main() {
	dsn := "root:aaa@tcp(calcnode.site:3306)/vrapi?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(dsn)
	database, _ := util.NewDatabase("", dsn)
	database.InitializeTables()
}
