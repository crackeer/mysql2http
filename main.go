package main

import (
	"github.com/crackeer/mysql2http/database"
	"github.com/crackeer/mysql2http/generator"
)

var databases map[string]string = map[string]string{
	"mysql2http": "root:1234567@tcp(127.0.0.1:3306)/mysql2http?charset=utf8&parseTime=True&loc=Local",
}

func main() {
	generator, err := generator.NewGoFileGenerator("./tmp")
	if err != nil {
		panic(err.Error())
	}
	mainData := []map[string]interface{}{}
	for name, dsn := range databases {
		database, _ := database.NewDatabase(name, dsn)
		if err := database.Initialize(); err != nil {
			panic(err.Error())
		}
		if err := generator.GenModelRouter(database.Name, dsn, database.BatchGenModelInput()); err != nil {
			panic(err.Error())
		}
		mainData = append(mainData, map[string]interface{}{
			"database": name,
			"tables":   database.GenMainRouterInput(),
		})
	}
	generator.CopySomeFiles()
	generator.GenMainGOFile(mainData)

}
