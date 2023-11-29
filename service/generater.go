package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/crackeer/mysql2http/service/template"
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

func regularGoFileName(table string) string {
	return strings.Trim(table, "_- ")
}

// GenerateRouter
//
//	@receiver g
//	@param dbName
//	@param dsn
//	@param tableFields
//	@return error
func (g *GoFileGenerator) GenModelRouter(dbName, dsn string, tableFields map[string]map[string]interface{}) error {
	bar := NewProgressBar(len(tableFields), "generating routers")
	for table, data := range tableFields {
		bar.Add(1)
		bytes, err := template.Render("database/table.tpl", data)
		if err != nil {
			return fmt.Errorf("render template table.tpl error: %v, database = %v | table = %s", err, dbName, table)
		}
		if err := g.write(filepath.Join("database", dbName, regularGoFileName(table)+".go"), bytes); err != nil {
			return fmt.Errorf("generate table router %v error %v", table, err)
		}
	}
	return nil
}

func (g *GoFileGenerator) GenContainer(list []map[string]interface{}) error {
	bytes, err := template.Render("container/database.tpl", map[string]interface{}{
		"databases": list,
	})
	if err != nil {
		return err
	}
	if err := g.write(filepath.Join("container/database.go"), bytes); err != nil {
		return fmt.Errorf("generate main.go error %v", err)
	}
	return nil
}

func (g *GoFileGenerator) GenMainGOFile(list []map[string]interface{}, port int) error {
	bytes, err := template.Render("main.tpl", map[string]interface{}{
		"databases": list,
		"port":      port,
	})
	if err != nil {
		return err
	}
	if err := g.write(filepath.Join("main.go"), bytes); err != nil {
		return fmt.Errorf("generate main.go error %v", err)
	}
	return nil
}

func (g *GoFileGenerator) CopySomeFiles() error {
	allFiles := template.ReadAllFileList()
	for _, item := range allFiles {
		if strings.HasSuffix(item, ".go") || strings.HasSuffix(item, ".mod") {
			bytes, err := template.Read(item)
			if err != nil {
				continue
			}
			if err := g.write(item, bytes); err != nil {
				return fmt.Errorf("copy `%s` error %v", item, err)
			}
		}
	}
	if bytes, err := template.Read("go_mod.tpl"); err == nil {
		if err := g.write("go.mod", bytes); err != nil {
			return fmt.Errorf("copy go_mod.tpl error %v", err)
		}
	}
	return nil
}
