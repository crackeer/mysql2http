package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/crackeer/mysql2http/database"
	"github.com/crackeer/mysql2http/generator"
)

type Config struct {
	Debug    bool `json:"debug"`
	Database []struct {
		Name string `json:"name"`
		DSN  string `json:"dsn"`
	} `json:"database"`
	CodeFolder string `json:"code_folder"`
}

var databases map[string]string = map[string]string{
	"mysql2http": "root:1234567@tcp(127.0.0.1:3306)/mysql2http?charset=utf8&parseTime=True&loc=Local",
}

func parseConfig() (*Config, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("no config file")
	}
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		return nil, err
	}
	cnf := &Config{}
	if err := json.Unmarshal(bytes, cnf); err != nil {
		return nil, fmt.Errorf("unmarshal config error: %v", err)
	}

	return cnf, nil

}

func main() {
	cnf, err := parseConfig()
	if err != nil {
		panic(err)
	}

	generator, err := generator.NewGoFileGenerator(cnf.CodeFolder)
	if err != nil {
		panic(err.Error())
	}
	if !cnf.Debug {
		defer os.RemoveAll(cnf.CodeFolder)
	}
	mainData := []map[string]interface{}{}
	for _, item := range cnf.Database {
		database, _ := database.NewDatabase(item.Name, item.DSN)
		if err := database.Initialize(); err != nil {
			panic(err.Error())
		}
		if err := generator.GenModelRouter(database.Name, item.DSN, database.BatchGenModelInput()); err != nil {
			panic(err.Error())
		}
		mainData = append(mainData, map[string]interface{}{
			"database": item.Name,
			"tables":   database.GenMainRouterInput(),
		})
	}
	generator.CopySomeFiles()
	generator.GenMainGOFile(mainData)

}
