package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// GoFileGenerator
type GoFileGenerator struct {
	WorkDir string
}

// NewGoFileGenerator
//
//	@param workDir
//	@return *GoFileGenerator
//	@return error
func NewGoFileGenerator(workDir string) (*GoFileGenerator, error) {
	if err := os.MkdirAll(workDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("create work_dir `%s` error: %v", workDir, err)
	}

	return &GoFileGenerator{
		WorkDir: workDir,
	}, nil
}

func (g *GoFileGenerator) write(file string, data []byte) error {
	fullPath := filepath.Join(g.WorkDir, file)
	parentDir, _ := filepath.Split(fullPath)
	if _, err := os.Stat(parentDir); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(parentDir, os.FileMode(0755)); err != nil {
			return fmt.Errorf("create directory `%s` failed: %v", parentDir, err)
		}
	}
	return os.WriteFile(fullPath, data, os.ModePerm)
}

// GenerateRouter
//
//	@receiver g
//	@param dbName
//	@param dsn
//	@param tableFields
//	@return error
func (g *GoFileGenerator) GenModelRouter(dbName, dsn string, tableFields map[string]map[string]interface{}) error {
	for table, data := range tableFields {
		bytes, _ := json.Marshal(data)
		fmt.Println(string(bytes))
		bytes, err := Render("model.tpl", data)
		if err != nil {
			return fmt.Errorf("render template model.tpl error: %v, database = %v | table = %s", err, dbName, table)
		}
		if err := g.write(filepath.Join("database", dbName, table+".go"), bytes); err != nil {
			return fmt.Errorf("generate table router %v error %v", table, err)
		}
	}
	bytes, err := Render("database.tpl", map[string]interface{}{
		"dsn":      dsn,
		"database": dbName,
	})
	if err != nil {
		return err
	}
	if err := g.write(filepath.Join("database", dbName, "db.go"), bytes); err != nil {
		return fmt.Errorf("generate %s / db.go error %v", dbName, err)
	}

	return nil
}

func (g *GoFileGenerator) GenMainGOFile(list []map[string]interface{}) error {
	bytes, err := Render("main.tpl", map[string]interface{}{
		"databases": list,
	})
	if err != nil {
		return err
	}
	if err := g.write(filepath.Join("main.go"), bytes); err != nil {
		return fmt.Errorf("generate main.go error %v", err)
	}
	return nil
}
