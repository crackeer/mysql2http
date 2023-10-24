package util

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

	Name       string
	DSN        string                  `json:"dsn"`
	TableField map[string][]TableField `json:"tables"`
}

func NewDatabase(name string, dsn string) (*Database, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Database{
		Name:       name,
		gormDB:     db,
		DSN:        dsn,
		TableField: map[string][]TableField{},
	}, nil
}

func (db *Database) Initialize() error {
	for _, table := range db.Tables() {
		fields, _ := db.Desc(table)
		db.TableField[table] = fields
	}
	return nil
}

func (db *Database) Close() error {
	return db.Close()
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

// BatchGenModelInput
//
//	@receiver db
//	@return map
func (db *Database) BatchGenModelInput() map[string]map[string]interface{} {
	tables := map[string]map[string]interface{}{}
	for table, fields := range db.TableField {
		tables[table] = map[string]interface{}{
			"database":          db.Name,
			"table":             table,
			"table_struct_name": strutil.UpperFirst(strutil.Camel(table)),
			"fields":            convertFields(fields),
		}
	}
	return tables
}

func (db *Database) GenMainRouterInput() []map[string]interface{} {
	tables := []map[string]interface{}{}
	for table := range db.TableField {
		tables = append(tables, map[string]interface{}{
			"table":             table,
			"table_struct_name": strutil.UpperFirst(strutil.Camel(table)),
		})
	}
	return tables
}

var typeTypeMapping map[string]string = map[string]string{
	"int":      "int64",
	"varchar":  "string",
	"datetime": "time.Time",
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
