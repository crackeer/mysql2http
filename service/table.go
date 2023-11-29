package service

import (
	"strings"

	"github.com/gookit/goutil/strutil"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TableField
type TableField struct {
	Field   string `gorm:"Field"`
	Type    string `gorm:"Type"`
	Null    string `gorm:"Null"`
	Key     string `gorm:"Key"`
	Default string `gorm:"Default"`
	Extra   string `gorm:"extra"`
}

// Database
type Database struct {
	gormDB *gorm.DB

	Name        string
	DSN         string                  `json:"dsn"`
	TableField  map[string][]TableField `json:"tables"`
	CreateTable map[string]string
}

func NewDatabase(name string, dsn string) (*Database, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Database{
		Name:        name,
		gormDB:      db,
		DSN:         dsn,
		TableField:  map[string][]TableField{},
		CreateTable: map[string]string{},
	}, nil
}

// Initialize
//
//	@receiver db
//	@return error
func (db *Database) Initialize() error {
	tables := db.Tables()
	bar := NewProgressBar(len(tables), "analyzing database tables")
	for _, table := range tables {
		fields, _ := db.Desc(table)
		db.TableField[table] = fields
		db.CreateTable[table], _ = db.ShowCreateTable(table)
		bar.Add(1)
	}
	return nil
}

// Tables
//
//	@param db
//	@return []string
func (db *Database) Tables() []string {
	list := []map[string]interface{}{}
	retData := []string{}
	if err := db.gormDB.Raw("show tables").Scan(&list).Error; err != nil {
		return retData
	}
	for _, value := range list {
		for _, value := range value {
			retData = append(retData, value.(string))
		}
	}
	return retData
}

// Desc
//
//	@param db
//	@param table
//	@return interface{}
//	@return error
func (db *Database) Desc(table string) ([]TableField, error) {
	list := []TableField{}

	if err := db.gormDB.Raw("desc " + table).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

// ShowCreateTable
//
//	@receiver db
//	@param table
//	@return string
//	@return error
func (db *Database) ShowCreateTable(table string) (string, error) {
	data := map[string]interface{}{}
	if err := db.gormDB.Raw("show create table  " + table).Scan(&data).Error; err != nil {
		return "", err
	}

	if value, ok := data["Create Table"]; ok {
		if stringValue, ok := value.(string); ok {
			return stringValue, nil
		}
	}

	return "", nil
}

// BatchGenModelData ...
//
//	@receiver db
//	@return map
func (db *Database) BatchGenModelData() map[string]map[string]interface{} {
	tables := map[string]map[string]interface{}{}
	for table, fields := range db.TableField {
		tables[table] = map[string]interface{}{
			"database":          db.Name,
			"table":             table,
			"table_struct_name": strutil.UpperFirst(strutil.Camel(table)),
			"fields":            convertFields(fields),
			"include_time":      includeDateTime(fields),
		}
	}
	return tables
}

// GenJSONData
//
//	@receiver db
//	@return map
func (db *Database) GenJSONData() map[string]interface{} {
	retData := map[string]interface{}{
		"database": db.Name,
	}
	tables := []map[string]interface{}{}
	for table, fields := range db.TableField {
		tables = append(tables, map[string]interface{}{
			"table":      table,
			"fields":     fields,
			"create_sql": db.CreateTable[table],
		})
	}
	retData["table"] = tables
	return retData
}

var typeTypeMapping map[string]string = map[string]string{
	"int":      "int64",
	"float":    "float64",
	"datetime": "define.LocalTime",
}

func getDataType(sqlType string) string {
	for key, value := range typeTypeMapping {
		if strings.Contains(sqlType, key) {
			return value
		}
	}
	return "string"
}

func convertFields(fields []TableField) []map[string]interface{} {
	retData := []map[string]interface{}{}
	for _, item := range fields {
		retData = append(retData, map[string]interface{}{
			"field": item.Field,
			"type":  getDataType(item.Type),
			"name":  strutil.UpperFirst(strutil.Camel(item.Field)),
		})
	}
	return retData
}

func includeDateTime(fields []TableField) bool {
	for _, item := range fields {
		if getDataType(item.Type) == "define.LocalTime" {
			return true
		}
	}
	return false
}
