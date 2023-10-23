package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database
type Database struct {
	gormDB *gorm.DB

	Name      string
	DSN       string                  `json:"dsn"`
	TableList map[string][]TableField `json:"tables"`
}

func NewDatabase(name string, dsn string) (*Database, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Database{
		Name:   name,
		gormDB: db,
		DSN:    dsn,
	}, nil
}

func (db *Database) InitializeTables() error {
	for _, table := range db.Tables() {
		fields, _ := db.Desc(table)
		db.TableList[table] = fields
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

// TableField
type TableField struct {
	Field   string `gorm:"Field"`
	Type    string `gorm:"Type"`
	Null    string `gorm:"Null"`
	Key     string `gorm:"Key"`
	Default string `gorm:"Default"`
	Extra   string `gorm:"extra"`
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
