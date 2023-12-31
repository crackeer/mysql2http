package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/crackeer/mysql2http/service"
	"github.com/logrusorgru/aurora/v4"
)

type Config struct {
	Debug    bool   `json:"debug"`
	Target   string `json:"target"`
	Database []struct {
		Name string `json:"name"`
		DSN  string `json:"dsn"`
	} `json:"database"`
	Port       int    `json:"port"`
	CodeFolder string `json:"code_folder"`
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
	os.RemoveAll(cnf.CodeFolder)

	generator, err := service.NewGoFileGenerator(cnf.CodeFolder)
	if err != nil {
		panic(err.Error())
	}

	mainData := []map[string]interface{}{}
	databaseData := []map[string]interface{}{}
	// ---> 生成数据库model文件
	for index, item := range cnf.Database {
		fmt.Printf("> %d. %s->%s\n", index+1, aurora.Green(item.Name), aurora.Green(item.DSN))
		database, err := service.NewDatabase(item.Name, item.DSN)
		if err != nil {
			panic(err.Error())
		}
		if err := database.Initialize(); err != nil {
			panic(fmt.Sprintf("parse table failed: %v", err))
		}
		if err := generator.GenModel(database.Name, item.DSN, database.BatchGenModelData()); err != nil {
			panic(fmt.Sprintf("generate router failed: %v[db = %s]", err, item.Name))
		}
		mainData = append(mainData, map[string]interface{}{
			"database": item.Name,
			"dsn":      item.DSN,
		})
		databaseData = append(databaseData, database.GenJSONData())
	}
	// -> 复制不需要生成的文件（go/html/js/css/json/go.mod）
	generator.CopyOriginFiles()
	// ---> 生成container下的go文件
	generator.GenContainer(mainData)
	// ---> 生成main.go
	generator.GenMainGOFile(mainData, cnf.Port)

	generator.GenStaticJSON(map[string]interface{}{
		"time":     time.Now().Unix(),
		"database": databaseData,
	})

	if len(cnf.Target) < 1 {
		fmt.Println(aurora.BrightYellow("We finished, the output code folder is: " + cnf.CodeFolder))
		return
	}

	if !cnf.Debug {
		defer os.RemoveAll(cnf.CodeFolder)
	}

	compiler := service.NewCompiler(cnf.CodeFolder)
	fmt.Println("Compiling...")
	if err := compiler.Prepare(); err != nil {
		panic(fmt.Sprintf("go mod tidy error: %v", err))
	}
	target, err := filepath.Abs(cnf.Target)
	if err != nil {
		panic(fmt.Sprintf("get output file error: %v", err))
	}

	if err := compiler.Build(target); err != nil {
		panic(fmt.Sprintf("build error: %v", err))
	}
	fmt.Println(aurora.BrightYellow("We finished, the output file is: " + target))
}
