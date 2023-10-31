package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

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
	if !cnf.Debug {
		defer os.RemoveAll(cnf.CodeFolder)
	}
	mainData := []map[string]interface{}{}
	for index, item := range cnf.Database {
		fmt.Printf("> %d. %s->%s\n", index+1, aurora.Green(item.Name), aurora.Green(item.DSN))
		database, err := service.NewDatabase(item.Name, item.DSN)
		if err != nil {
			panic(err.Error())
		}
		if err := database.Initialize(); err != nil {
			panic(fmt.Sprintf("parse table failed: %v", err))
		}
		if err := generator.GenModelRouter(database.Name, item.DSN, database.BatchGenModelInput()); err != nil {
			panic(fmt.Sprintf("generate router failed: %v[db = %s]", err, item.Name))
		}
		mainData = append(mainData, map[string]interface{}{
			"database": item.Name,
			"tables":   database.GenMainRouterInput(),
		})
	}
	generator.CopySomeFiles()
	generator.GenMainGOFile(mainData, cnf.Port)

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
